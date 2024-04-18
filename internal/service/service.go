package service

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/rbf"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/snippets"
	"github.com/gernest/requiemdb/internal/store"
	"go.opentelemetry.io/otel/metric"
	collector_logs "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	collector_metrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	collector_trace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

const (
	// Size of buffered channel that accepts samples.
	DataBuffer = 4 << 10
)

type Service struct {
	snippets *snippets.Snippets
	store    *store.Storage
	dataMu   sync.Mutex
	data     *samples.List
	seq      *seq.Seq
	stats    struct {
		processed metric.Int64Counter
	}
	v1.UnsafeRQServer
}

func NewService(ctx context.Context, db *badger.DB,
	seq *seq.Seq,
	idx *rbf.DB,
	listen string, retention time.Duration) (*Service, error) {
	sn, err := snippets.New()
	if err != nil {
		return nil, err
	}
	tree, err := lsm.New(db, seq)
	if err != nil {
		return nil, err
	}
	storage, err := store.NewStore(db, idx, seq, tree)
	if err != nil {
		return nil, err
	}

	service := &Service{
		snippets: sn,
		store:    storage,
		seq:      seq,
		data:     samples.Get(),
	}
	service.stats.processed, err = self.Meter().Int64Counter(
		"samples.processed",
		metric.WithDescription("Total number of samples processed"),
	)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *Service) Start(ctx context.Context) {
	go s.store.Start(ctx)
	go s.start(ctx)
}

func (s *Service) start(ctx context.Context) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			s.dataMu.Lock()
			ls := s.data
			s.data = samples.Get()
			s.dataMu.Unlock()
			err := s.store.SaveSamples(ls)
			if err != nil {
				slog.Error("Failed to save samples", "err", err)
			}
			ls.Release()
		}
	}
}

func (s *Service) Save(data *v1.Data) error {
	s.dataMu.Lock()
	s.data.Items = append(s.data.Items, &v1.Sample{
		Id:   s.seq.SampleID(),
		Data: data,
	})
	return nil
}

func (s *Service) Close() {
	s.store.Close()
	s.snippets.Close()
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
	err := r.svc.Save(&v1.Data{
		Data: &v1.Data_Metrics{Metrics: &metricsv1.MetricsData{
			ResourceMetrics: req.ResourceMetrics,
		}},
	})
	if err != nil {
		slog.Error("failed saving data", "err", err)
		return nil, err
	}
	return &collector_metrics.ExportMetricsServiceResponse{}, nil
}

type Logs struct {
	collector_logs.UnsafeLogsServiceServer
	svc *Service
}

var _ collector_logs.LogsServiceServer = (*Logs)(nil)

func (r *Logs) Export(ctx context.Context, req *collector_logs.ExportLogsServiceRequest) (*collector_logs.ExportLogsServiceResponse, error) {
	err := r.svc.Save(&v1.Data{
		Data: &v1.Data_Logs{
			Logs: &logsv1.LogsData{
				ResourceLogs: req.ResourceLogs,
			},
		},
	})
	if err != nil {
		slog.Error("failed saving data", "err", err)
		return nil, err
	}
	return &collector_logs.ExportLogsServiceResponse{}, nil
}

type Trace struct {
	collector_trace.UnsafeTraceServiceServer
	svc *Service
}

var _ collector_trace.TraceServiceServer = (*Trace)(nil)

func (r *Trace) Export(ctx context.Context, req *collector_trace.ExportTraceServiceRequest) (*collector_trace.ExportTraceServiceResponse, error) {
	err := r.svc.Save(&v1.Data{
		Data: &v1.Data_Traces{
			Traces: &tracev1.TracesData{
				ResourceSpans: req.ResourceSpans,
			},
		},
	})
	if err != nil {
		slog.Error("failed saving data", "err", err)
		return nil, err
	}
	return &collector_trace.ExportTraceServiceResponse{}, nil
}
