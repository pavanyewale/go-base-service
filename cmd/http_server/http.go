package http_server

import (
	"context"
	"fmt"
	"gobaseservice/internal/configs"
	"gobaseservice/internal/repository"
	"gobaseservice/internal/service"
	"gobaseservice/pkg/api/gobaseservice"
	"gobaseservice/pkg/utils"
	"net/http"

	"github.com/gofreego/goutils/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HTTPServer struct {
	cfg    *configs.Configuration
	server *http.Server
}

func (a *HTTPServer) Name() string {
	return "HTTP_Server"
}

func (a *HTTPServer) Shutdown(ctx context.Context) {
	if err := a.server.Shutdown(ctx); err != nil {
		logger.Panic(ctx, "failed to shutdown %s : %v", a.Name(), err)
	}
}

func NewHTTPServer(cfg *configs.Configuration) *HTTPServer {
	return &HTTPServer{
		cfg: cfg,
	}
}

func (a *HTTPServer) Run(ctx context.Context) {

	if a.cfg.Server.HTTPPort == 0 {
		logger.Panic(ctx, "http port is not provided")
	}

	repository := repository.NewRepository(ctx, &a.cfg.Repository)

	serviceSf := service.NewServiceFactory(ctx, &a.cfg.Service, repository)

	mux := runtime.NewServeMux()
	utils.RegisterSwaggerHandler(ctx, mux, "/swagger", "./docs/proto/", "/v1/gobaseservice.swagger.json")
	err := gobaseservice.RegisterBaseServiceHandlerServer(ctx, mux, serviceSf.PingService)
	if err != nil {
		logger.Panic(ctx, "failed to register ping service : %v", err)
	}
	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Server.HTTPPort),
		Handler: mux,
	}

	logger.Info(ctx, "Starting HTTP server on port %d", a.cfg.Server.HTTPPort)
	logger.Info(ctx, "Swagger UI is available at `http://localhost:%d/swagger`", a.cfg.Server.HTTPPort)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Panic(ctx, "failed to start http server : %v", err)
	}
}
