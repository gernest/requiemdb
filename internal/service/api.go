package service

import (
	"context"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/version"
)

var _ v1.RQServer = (*Service)(nil)

func (*Service) GetVersion(_ context.Context, _ *v1.GetVersionRequest) (*v1.Version, error) {
	return &v1.Version{
		Version: version.VERSION,
	}, nil
}
