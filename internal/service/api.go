package service

import (
	"context"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/version"
)

var _ v1.RQServer = (*Service)(nil)

func (s *Service) ScanSamples(ctx context.Context, req *v1.Scan) (*v1.Data, error) {
	return s.store.Scan(ctx, req)
}

func (*Service) GetVersion(_ context.Context, _ *v1.GetVersionRequest) (*v1.Version, error) {
	return &v1.Version{Version: version.VERSION}, nil
}
