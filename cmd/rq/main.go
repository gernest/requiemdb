package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
	"github.com/requiemdb/requiemdb/internal/logger"
	"github.com/requiemdb/requiemdb/internal/service"
	"github.com/urfave/cli/v3"
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
				Name:    "otlp-listen",
				Value:   ":4317",
				Usage:   "host:port address to listen to otlp collector grpc service",
				Sources: cli.EnvVars("RQ_OTLP_LISTEN"),
			},
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
	api, err := service.NewService(ctx, db, lsn)
	if err != nil {
		return err
	}
	defer api.Close()

	svr := &http.Server{
		Addr:        lsn,
		Handler:     api,
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go func() {
		defer cancel()
		slog.Info("starting http server", "address", lsn)
		exit = svr.ListenAndServe()
	}()
	<-ctx.Done()
	return
}

func setupLogging(cmd *cli.Command, o badger.Options) badger.Options {
	var level slog.Level
	level.UnmarshalText([]byte(cmd.String("logLevel")))
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: level,
				},
			),
		),
	)
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
