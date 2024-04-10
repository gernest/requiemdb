package query

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/commands"
	"github.com/gernest/requiemdb/internal/compile"
	"github.com/gernest/requiemdb/internal/js"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/render"
	"github.com/urfave/cli/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const doc = `# executes a .js or .ts file 
rq query script.ts

# or
rq query script.js

# only @requiemdb/rq package can be imported by a script
# if there is a package you can't live without please
# open a feature request on github`

func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "query",
		Usage:       "executes a js scripts that queries and process samples",
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
	file := cmd.Args().First()
	if file == "" {
		return errors.New("missing .js or .ts file to execute")
	}
	logger.Setup(cmd.Root().String("logLevel"))
	log := slog.Default().With("file", file)
	log.Debug("opening remote connection", "target", cmd.String("hostPort"))
	conn, err := grpc.Dial(cmd.String("hostPort"), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		return err
	}
	defer conn.Close()

	rq := v1.NewRQClient(conn)
	log.Debug("reading script")
	data, err := os.ReadFile(cmd.Args().First())
	if err != nil {
		return err
	}
	log.Debug("Compiling script")
	compiled, err := build(log, data)
	if err != nil {
		return err
	}
	log.Debug("executing script")

	vm := js.New().
		WithScan(func(s *v1.Scan) (*v1.Data, error) {
			return rq.ScanSamples(ctx, s)
		}).
		WithNow(time.Now)
	defer vm.Release()
	err = vm.Run(compiled)
	if err != nil {
		return err
	}
	if cmd.Bool("logs") {
		os.Stdout.Write(vm.Log.Bytes())
	}
	return render.Result(os.Stdout, vm.Output)
}

func build(log *slog.Logger, data []byte) (*goja.Program, error) {
	hash := xxhash.Sum64(data)
	key := filepath.Join(commands.Cache(), fmt.Sprint(hash))
	cached, err := os.ReadFile(key)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug("cache miss, compiling fresh package")
			cached, err = compile.Compile(data)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(key, cached, 0600)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		log.Debug("using cached package", "key", key)
	}
	return goja.Compile("scan.js", string(cached), true)
}
