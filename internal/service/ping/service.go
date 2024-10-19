package ping

import (
	"context"
	"gobaseservice/pkg/api/gobaseservice"

	"github.com/gofreego/goutils/logger"
)

type Config struct {
}

type Repository interface {
	Ping(ctx context.Context) error
}

type Service struct {
	cfg  *Config
	repo Repository
	gobaseservice.UnimplementedBaseServiceServer
}

// Ping implements ping.PingServiceServer.
func (s *Service) Ping(ctx context.Context, req *gobaseservice.PingRequest) (*gobaseservice.PingResponse, error) {
	logger.Debug(ctx, "Ping request received, %v", req.Message)
	err := s.repo.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &gobaseservice.PingResponse{
		Message: "Pong",
	}, nil
}

func NewService(ctx context.Context, cfg *Config, repo Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}
