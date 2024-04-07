package service

import (
	"context"
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/js"
	"github.com/gernest/requiemdb/internal/version"
)

var _ v1.RQServer = (*Service)(nil)

func (s *Service) Query(_ context.Context, req *v1.QueryRequest) (*v1.QueryResponse, error) {
	vm := js.New().
		WithScan(s.store.Scan).
		WithNow(time.Now)
	defer vm.Release()
	program, err := s.snippets.GetProgramData(req.Query)
	if err != nil {
		return nil, err
	}
	err = vm.Run(program)
	if err != nil {
		return nil, err
	}
	res := &v1.QueryResponse{
		Result: vm.Output,
	}
	if req.IncludeLogs {
		res.Logs = vm.Log.Bytes()
	}
	return res, nil
}

func (*Service) GetVersion(_ context.Context, _ *v1.GetVersionRequest) (*v1.Version, error) {
	return &v1.Version{
		Version: version.VERSION,
	}, nil
}
