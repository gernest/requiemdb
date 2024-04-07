package service

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/snippets"
	"github.com/gernest/requiemdb/internal/store"
	"github.com/gernest/requiemdb/ui"
	"github.com/go-chi/cors"
	grpc_protovalidate "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric"
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

var (
	rootFS, _  = fs.Sub(ui.FS, "dist")
	fileServer = http.FileServer(http.FS(rootFS))
)

type Service struct {
	db       *badger.DB
	snippets *snippets.Snippets
	store    *store.Storage
	hand     http.Handler
	data     chan *v1.Data

	stats struct {
		processed metric.Int64Counter
	}
	v1.UnsafeRQServer
}

func NewService(ctx context.Context, db *badger.DB, listen string, retention time.Duration) (*Service, error) {
	valid, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	sn, err := snippets.New()
	if err != nil {
		return nil, err
	}
	tree, err := lsm.New(db)
	if err != nil {
		return nil, err
	}
	storage, err := store.NewStore(db, tree)
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

	service := &Service{
		db:       db,
		snippets: sn,
		store:    storage,
		data:     make(chan *v1.Data, 4<<10),
	}
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
	service.stats.processed, err = self.Meter().Int64Counter(
		"samples.processed",
		metric.WithDescription("Total number of samples processed"),
	)
	if err != nil {
		service.Close()
		return nil, err
	}
	return service, nil
}

func (s *Service) Start(ctx context.Context) {
	go s.store.Start(ctx)
	go s.save(ctx)
}

func (s *Service) save(ctx context.Context) {
	for data := range s.data {
		err := s.store.Save(data)
		if err != nil {
			slog.Error("failed saving data sample", "err", err)
		}
		s.stats.processed.Add(ctx, 1)
	}
}

func (s *Service) Close() {
	close(s.data)
	s.store.Close()
	s.snippets.Close()
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
	return &Metrics{svc: s}
}
func (s *Service) Trace() *Trace {
	return &Trace{svc: s}
}

func (s *Service) Logs() *Logs {
	return &Logs{svc: s}
}

type Metrics struct {
	collector_metrics.UnsafeMetricsServiceServer
	svc *Service
}

var _ collector_metrics.MetricsServiceServer = (*Metrics)(nil)

func (r *Metrics) Export(ctx context.Context, req *collector_metrics.ExportMetricsServiceRequest) (*collector_metrics.ExportMetricsServiceResponse, error) {
	r.svc.data <- &v1.Data{
		Data: &v1.Data_Metrics{Metrics: &metricsv1.MetricsData{
			ResourceMetrics: req.ResourceMetrics,
		}},
	}
	return &collector_metrics.ExportMetricsServiceResponse{}, nil
}

type Logs struct {
	collector_logs.UnsafeLogsServiceServer
	svc *Service
}

var _ collector_logs.LogsServiceServer = (*Logs)(nil)

func (r *Logs) Export(ctx context.Context, req *collector_logs.ExportLogsServiceRequest) (*collector_logs.ExportLogsServiceResponse, error) {
	r.svc.data <- &v1.Data{
		Data: &v1.Data_Logs{
			Logs: &logsv1.LogsData{
				ResourceLogs: req.ResourceLogs,
			},
		},
	}
	return &collector_logs.ExportLogsServiceResponse{}, nil
}

type Trace struct {
	collector_trace.UnsafeTraceServiceServer
	svc *Service
}

var _ collector_trace.TraceServiceServer = (*Trace)(nil)

func (r *Trace) Export(ctx context.Context, req *collector_trace.ExportTraceServiceRequest) (*collector_trace.ExportTraceServiceResponse, error) {
	r.svc.data <- &v1.Data{
		Data: &v1.Data_Trace{
			Trace: &tracev1.TracesData{
				ResourceSpans: req.ResourceSpans,
			},
		},
	}
	return &collector_trace.ExportTraceServiceResponse{}, nil
}
