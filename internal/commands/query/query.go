package query

import (
	"context"
	"log/slog"
	"os"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/urfave/cli/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const doc = `rq query script.ts`

func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "query",
		Usage:       "sends query to rq instance",
		Description: doc,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "hostPort",
				Usage: "host:port address of rq",
				Value: "localhost:8080",
			},
			&cli.BoolFlag{
				Name:  "logs",
				Usage: "collects console.log output",
				Value: true,
			},
		},
		Action: run,
	}
}

func run(ctx context.Context, cmd *cli.Command) error {
	logger.Setup(cmd.Root().String("logLevel"))
	slog.Debug("opening remote connection", "target", cmd.String("hostPort"))
	conn, err := grpc.Dial(cmd.String("hostPort"), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		return err
	}
	defer conn.Close()

	rq := v1.NewRQClient(conn)
	slog.Debug("reading script file", "file", cmd.Args().First())
	data, err := os.ReadFile(cmd.Args().First())
	if err != nil {
		return err
	}
	slog.Debug("sending query request", "IncludeLogs", cmd.Bool("logs"))
	r, err := rq.Query(ctx, &v1.QueryRequest{
		Query:       data,
		IncludeLogs: cmd.Bool("logs"),
	})
	if err != nil {
		return err
	}
	if len(r.Logs) != 0 {
		os.Stdout.Write(r.Logs)
	}
	o, err := protojson.MarshalOptions{Multiline: true}.Marshal(r.Result)
	if err != nil {
		return err
	}
	os.Stdout.Write(o)
	return nil
}
