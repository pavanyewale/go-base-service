package grpc_server

import (
	"context"
	"fmt"
	"gobaseservice/internal/configs"
	"gobaseservice/internal/repository"
	"gobaseservice/internal/service"
	"gobaseservice/pkg/api/gobaseservice"
	"net"
	"net/http"

	"github.com/gofreego/goutils/logger"
	"google.golang.org/grpc"
	// Update
)

type GRPCServer struct {
	cfg    *configs.Configuration
	server *http.Server
}

func (a *GRPCServer) Name() string {
	return "GRPC_Server"
}

func (a *GRPCServer) Shutdown(ctx context.Context) {
	if err := a.server.Shutdown(ctx); err != nil {
		logger.Panic(ctx, "failed to shutdown %s : %v", a.Name(), err)
	}
}

func NewGRPCServer(cfg *configs.Configuration) *GRPCServer {
	return &GRPCServer{
		cfg: cfg,
	}
}

func (a *GRPCServer) Run(ctx context.Context) {

	if a.cfg.Server.GRPCPort == 0 {
		logger.Panic(ctx, "grpc port is not provided")
	}

	repository := repository.NewRepository(ctx, &a.cfg.Repository)

	serviceSf := service.NewServiceFactory(ctx, &a.cfg.Service, repository)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	gobaseservice.RegisterBaseServiceServer(grpcServer, serviceSf.PingService)

	logger.Info(ctx, "Starting gRPC server on port %d", a.cfg.Server.GRPCPort)

	// Listen on a TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Server.GRPCPort))
	if err != nil {
		logger.Panic(ctx, "failed to listen for grpc server: %v", err)
	}

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		logger.Panic(ctx, "failed to start grpc server: %v", err)
	}
}
