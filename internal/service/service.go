package service

import (
	"context"
	"net/http"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-chi/cors"
	grpc_protovalidate "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/snippets"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	db       *badger.DB
	snippets *snippets.Snippets
	hand     http.Handler
	v1.UnimplementedRQServer
}

func NewService(ctx context.Context, db *badger.DB, listen string) (*Service, error) {
	valid, err := protovalidate.New()
	if err != nil {
		return nil, err
	}
	sn, err := snippets.New(db, 0)
	if err != nil {
		return nil, err
	}
	svr := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_protovalidate.StreamServerInterceptor(valid),
		),
		grpc.UnaryInterceptor(
			grpc_protovalidate.UnaryServerInterceptor(valid),
		),
	)

	service := &Service{db: db, snippets: sn}
	v1.RegisterRQServer(svr, service)
	web := grpcweb.WrapServer(svr,
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}))
	api := runtime.NewServeMux()
	reflection.Register(svr)
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = v1.RegisterRQHandlerFromEndpoint(
		ctx, api, listen, dopts,
	)
	if err != nil {
		return nil, err
	}

	base := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index.html", "/logo.svg", "/robot.txt":
			// fileServer.ServeHTTP(w, r)
			return
		default:
			if strings.HasPrefix(r.URL.Path, "/api/v1/") {
				api.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/static/") {
				// fileServer.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	root := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			web.ServeHTTP(w, r)
			return
		}
		base.ServeHTTP(w, r)
	}), &http2.Server{})
	service.hand = corsMiddleware().Handler(root)
	return service, nil
}

func (s *Service) Close() error {
	return nil
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.hand.ServeHTTP(w, r)
}

func corsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowCredentials: true,
	})
}
