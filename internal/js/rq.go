package js

import (
	"errors"
	"fmt"
	"io"

	"github.com/dop251/goja"
	"github.com/requiemdb/requiemdb/internal/logger"
	"github.com/requiemdb/requiemdb/packages/rq"
)

var ErrInvalidModule = errors.New("js: invalid module")

var program *goja.Program

func init() {
	var err error
	program, err = goja.Compile("rq", rq.RQ, true)
	if err != nil {
		logger.Fail("failed compiling rq package", "err", err)
	}
}

func New(out io.Writer) (*goja.Runtime, error) {
	r := goja.New()
	err := r.Set("require", requireModule(r))
	if err != nil {
		return nil, err
	}
	return r, r.Set("console", console(r, out))
}

func requireModule(r *goja.Runtime) func(call *goja.FunctionCall) goja.Value {
	return func(call *goja.FunctionCall) goja.Value {
		name := call.Arguments[0].String()
		if name != "rq" {
			panic(fmt.Errorf("module %q is not found", name))
		}
		_, err := r.RunProgram(program)
		if err != nil {
			panic(err)
		}
		return r.GlobalObject().Get("exports")
	}
}
