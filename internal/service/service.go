package service

import (
	"context"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-chi/cors"
	grpc_protovalidate "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/lsm"
	"github.com/requiemdb/requiemdb/internal/snippets"
	"github.com/requiemdb/requiemdb/internal/store"
	"github.com/requiemdb/requiemdb/internal/transform"
	"github.com/requiemdb/requiemdb/ui"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	collector_logs "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	collector_metrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	collector_trace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var rootFS, _ = fs.Sub(ui.FS, "dist")

var fileServer = http.FileServer(http.FS(rootFS))

type Service struct {
	db        *badger.DB
	snippets  *snippets.Snippets
	seq       *badger.Sequence
	tree      *lsm.Tree
	retention time.Duration
	hand      http.Handler
	v1.UnimplementedRQServer
}

func NewService(ctx context.Context, db *badger.DB, seq *badger.Sequence, listen string, retention time.Duration) (*Service, error) {
	valid, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	sn, err := snippets.New(db, 0)
	if err != nil {
		return nil, err
	}
	tree, err := lsm.New(db)
	if err != nil {
		return nil, err
	}
	svr := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_protovalidate.StreamServerInterceptor(valid),
		),
		grpc.UnaryInterceptor(
			grpc_protovalidate.UnaryServerInterceptor(valid),
		),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	service := &Service{db: db, snippets: sn, seq: seq, tree: tree, retention: retention}
	v1.RegisterRQServer(svr, service)
	web := grpcweb.WrapServer(svr,
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}))
	api := runtime.NewServeMux()
	reflection.Register(svr)
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = v1.RegisterRQHandlerFromEndpoint(
		ctx, api, listen, dopts,
	)
	if err != nil {
		return nil, err
	}

	base := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index.html", "/logo.svg", "/robot.txt":
			fileServer.ServeHTTP(w, r)
			return
		default:
			if strings.HasPrefix(r.URL.Path, "/api/v1/") {
				api.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/assets/") {
				fileServer.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	root := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			web.ServeHTTP(w, r)
			return
		}
		base.ServeHTTP(w, r)
	}), &http2.Server{})
	service.hand = corsMiddleware().Handler(root)
	return service, nil
}

func (s *Service) Start(ctx context.Context) {
	go s.tree.Start(ctx)
}

func (s *Service) Close() error {
	return nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.hand.ServeHTTP(w, r)
}

func corsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowCredentials: true,
	})
}

func (s *Service) Metrics() *Metrics {
	return &Metrics{db: s.db, seq: s.seq, tree: s.tree, retention: s.retention}
}
func (s *Service) Trace() *Trace {
	return &Trace{db: s.db, seq: s.seq, tree: s.tree, retention: s.retention}
}

func (s *Service) Logs() *Logs {
	return &Logs{db: s.db, seq: s.seq, tree: s.tree, retention: s.retention}
}

type Metrics struct {
	collector_metrics.UnsafeMetricsServiceServer
	db        *badger.DB
	seq       *badger.Sequence
	tree      *lsm.Tree
	retention time.Duration
}

var _ collector_metrics.MetricsServiceServer = (*Metrics)(nil)

func (r *Metrics) Export(ctx context.Context, req *collector_metrics.ExportMetricsServiceRequest) (*collector_metrics.ExportMetricsServiceResponse, error) {
	o := transform.NewContext()
	defer o.Release()

	sample, labels, err := o.Process(&metricsv1.MetricsData{
		ResourceMetrics: req.ResourceMetrics,
	})
	if err != nil {
		return nil, err
	}
	defer labels.Release()
	err = store.Store(r.db, r.tree, r.seq, labels, sample, r.retention, v1.RESOURCE_METRICS)
	if err != nil {
		return nil, err
	}
	return &collector_metrics.ExportMetricsServiceResponse{}, nil
}

type Logs struct {
	collector_logs.UnsafeLogsServiceServer
	db        *badger.DB
	tree      *lsm.Tree
	seq       *badger.Sequence
	retention time.Duration
}

var _ collector_logs.LogsServiceServer = (*Logs)(nil)

func (r *Logs) Export(ctx context.Context, req *collector_logs.ExportLogsServiceRequest) (*collector_logs.ExportLogsServiceResponse, error) {
	o := transform.NewContext()
	defer o.Release()

	sample, labels, err := o.Process(&logsv1.LogsData{
		ResourceLogs: req.ResourceLogs,
	})
	if err != nil {
		return nil, err
	}
	defer labels.Release()
	err = store.Store(r.db, r.tree, r.seq, labels, sample, r.retention, v1.RESOURCE_LOGS)
	if err != nil {
		return nil, err
	}
	return &collector_logs.ExportLogsServiceResponse{}, nil
}

type Trace struct {
	collector_trace.UnsafeTraceServiceServer
	db        *badger.DB
	seq       *badger.Sequence
	tree      *lsm.Tree
	retention time.Duration
}

var _ collector_trace.TraceServiceServer = (*Trace)(nil)

func (r *Trace) Export(ctx context.Context, req *collector_trace.ExportTraceServiceRequest) (*collector_trace.ExportTraceServiceResponse, error) {
	o := transform.NewContext()
	defer o.Release()

	sample, labels, err := o.Process(&tracev1.TracesData{
		ResourceSpans: req.ResourceSpans,
	})
	if err != nil {
		return nil, err
	}
	defer labels.Release()
	err = store.Store(r.db, r.tree, r.seq, labels, sample, r.retention, v1.RESOURCE_TRACES)
	if err != nil {
		return nil, err
	}
	return &collector_trace.ExportTraceServiceResponse{}, nil
}
