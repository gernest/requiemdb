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

func Home(base ...string) string {
	if len(base) == 0 || len(base) == 1 && base[0] == "" {
		base = []string{home.Dir(), ".rq"}
	}
	h := filepath.Join(base...)
	os.MkdirAll(h, 0755)
	return h
}

func Cache(base ...string) string {
	h := filepath.Join(Home(base...), "cache")
	os.MkdirAll(h, 0755)
	return h
}

func Data(base ...string) string {
	h := filepath.Join(Home(base...), "data")
	os.MkdirAll(h, 0755)
	return h
}

func DB(base ...string) string {
	h := filepath.Join(Data(base...), "db")
	os.MkdirAll(h, 0755)
	return h
}

func Index(base ...string) string {
	h := filepath.Join(Data(base...), "index")
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
