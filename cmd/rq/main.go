package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
	"github.com/gernest/requiemdb/internal/commands/query"
	"github.com/gernest/requiemdb/internal/commands/version"
	"github.com/gernest/requiemdb/internal/gc"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/service"
	"github.com/urfave/cli/v3"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	collector_logs "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	collector_metrics "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	collector_trace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/grpc"
)

func main() {
	cmd := cli.Command{
		Name:  "rq",
		Usage: "OpenTelemetry database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "logLevel",
				Value:   "info",
				Sources: cli.EnvVars("RQ_LOG_LEVEL"),
			},
			&cli.StringFlag{
				Name:    "listen",
				Usage:   "HTTP address to bind api server",
				Value:   ":8080",
				Sources: cli.EnvVars("RQ_LISTEN"),
			},
			&cli.StringFlag{
				Name:    "otlpListen",
				Value:   ":4317",
				Usage:   "host:port address to listen to otlp collector grpc service",
				Sources: cli.EnvVars("RQ_OTLP_LISTEN"),
			},
			&cli.DurationFlag{
				Name:    "retentionPeriod",
				Value:   7 * 24 * time.Hour, //one week
				Sources: cli.EnvVars("RQ_DATA_RETENTION_PERIOD"),
			},
		},
		Commands: []*cli.Command{
			query.Cmd(),
			version.Cmd(),
		},
		Action: run,
	}
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		logger.Fail("exited server", "err", err)
	}
}

func run(ctx context.Context, cmd *cli.Command) (exit error) {
	data := cmd.Args().First()
	o := badger.DefaultOptions(data).
		WithLogger(Logger{}).
		WithCompression(
			options.ZSTD,
		)
	o = setupLogging(cmd, o)
	if data == "" {
		slog.Warn("missing data path, opening in memory database")
		o = o.WithInMemory(true)
	}
	db, err := badger.Open(o)
	if err != nil {
		return err
	}
	defer db.Close()

	lsn := cmd.String("listen")
	api, err := service.NewService(ctx, db, lsn, cmd.Duration("retentionPeriod"))
	if err != nil {
		return err
	}
	defer api.Close()

	svr := &http.Server{
		Addr:        lsn,
		Handler:     api,
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}

	otelAddress := cmd.String("otlpListen")
	otelGRPC, err := net.Listen("tcp", otelAddress)
	if err != nil {
		logger.Fail("failed listening otlp address", "addr", otelAddress, "err", err)
	}
	defer otelGRPC.Close()

	oSvr := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	collector_metrics.RegisterMetricsServiceServer(oSvr, api.Metrics())
	collector_logs.RegisterLogsServiceServer(oSvr, api.Logs())
	collector_trace.RegisterTraceServiceServer(oSvr, api.Trace())

	providers, err := self.Setup(ctx, api.Metrics(), api.Trace())
	if err != nil {
		return err
	}
	defer providers.Shutdown(context.Background())

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go api.Start(ctx)

	go func() {
		defer cancel()
		slog.Info("starting http server", "address", lsn)
		exit = svr.ListenAndServe()
	}()
	go func() {
		defer cancel()
		slog.Info("starting gRPC otel collector server", "address", otelAddress)
		err := oSvr.Serve(otelGRPC)
		if err != nil {
			slog.Error("exited grpc service", "err", err)
		}
	}()
	// start gc  routines
	gc.Run(ctx, db)
	<-ctx.Done()
	svr.Shutdown(context.Background())
	oSvr.GracefulStop()
	return
}

func setupLogging(cmd *cli.Command, o badger.Options) badger.Options {
	level := logger.Setup(cmd.String("logLevel"), os.Stdout)
	switch level {
	case slog.LevelInfo:
		return o.WithLoggingLevel(badger.INFO)
	case slog.LevelDebug:
		return o.WithLoggingLevel(badger.DEBUG)
	case slog.LevelWarn:
		return o.WithLoggingLevel(badger.WARNING)
	case slog.LevelError:
		return o.WithLoggingLevel(badger.ERROR)
	default:
		return o
	}
}

type Logger struct{}

var _ badger.Logger = (*Logger)(nil)

func (Logger) Errorf(msg string, args ...interface{}) {
	slog.Error(fmt.Sprintf(msg, args...))
}
func (Logger) Warningf(msg string, args ...interface{}) {
	slog.Warn(fmt.Sprintf(msg, args...))
}
func (Logger) Infof(msg string, args ...interface{}) {
	slog.Info(fmt.Sprintf(msg, args...))
}
func (Logger) Debugf(msg string, args ...interface{}) {
	slog.Debug(fmt.Sprintf(msg, args...))
}
