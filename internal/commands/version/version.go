package version

import (
	"context"
	"os"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/commands"
	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Prints remote rq instance version",
		Flags: commands.FLags(),
		Action: func(ctx context.Context, c *cli.Command) error {
			conn, err := commands.Conn(c)
			if err != nil {
				return err
			}
			defer conn.Close()
			v, err := v1.NewRQClient(conn).GetVersion(ctx, &v1.GetVersionRequest{})
			if err != nil {
				return err
			}
			os.Stdout.WriteString(v.Version)
			return nil
		},
	}
}
