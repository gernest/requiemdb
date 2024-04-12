package commands

import (
	"os"
	"path/filepath"

	_ "github.com/gernest/requiemdb/internal/compress"
	"github.com/gernest/requiemdb/internal/home"
	"github.com/urfave/cli/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Home() string {
	h := filepath.Join(home.Dir(), ".rq")
	os.MkdirAll(h, 0755)
	return h
}

func Cache() string {
	h := filepath.Join(Home(), "cache")
	os.MkdirAll(h, 0755)
	return h
}

func FLags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "remoteAddress",
			Usage:   "host:port address of remote rq",
			Value:   "localhost:4317",
			Sources: cli.EnvVars("RQ_REMOTE_ADDRESS"),
		},
	}
}

func Conn(cmd *cli.Command) (*grpc.ClientConn, error) {
	return grpc.Dial(cmd.String("remoteAddress"), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
}
