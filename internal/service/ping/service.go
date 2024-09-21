package ping

import (
	"context"
	"gobaseservice/pkg/api/v1/ping"
)

type Config struct {
}

type Repository interface {
	Ping(ctx context.Context) error
}

type Service struct {
	cfg  *Config
	repo Repository
	ping.UnimplementedPingServiceServer
}

// Ping implements ping.PingServiceServer.
func (s *Service) Ping(ctx context.Context, req *ping.PingRequest) (*ping.PingResponse, error) {
	err := s.repo.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &ping.PingResponse{}, nil
}

func NewService(ctx context.Context, cfg *Config, repo Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

// func (s *Service) Ping(ctx context.Context) error {
// }
