package service

import (
	"context"
	"errors"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/compress"
	"github.com/requiemdb/requiemdb/internal/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.RQServer = (*Service)(nil)

func (a *Service) GetSnippet(ctx context.Context, req *v1.GetSnippetRequest) (*v1.GetSnippetResponse, error) {
	res, err := a.snippets.Get(req.Name)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, status.Error(codes.NotFound, "Snippet not found")
		}
		return nil, err
	}
	data, err := compress.Decompress(res.Raw)
	if err != nil {
		return nil, err
	}
	return &v1.GetSnippetResponse{
		Raw: data,
	}, nil
}
func (*Service) GetVersion(_ context.Context, _ *v1.GetVersionRequest) (*v1.Version, error) {
	return &v1.Version{
		Version: version.VERSION,
	}, nil
}
