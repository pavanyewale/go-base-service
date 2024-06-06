package ping

import "context"

type Config struct {
}

type Repository interface {
	Ping(ctx context.Context) error
}

type Service struct {
	cfg  *Config
	repo Repository
}

func NewService(ctx context.Context, cfg *Config, repo Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) Ping(ctx context.Context) error {
	return s.repo.Ping(ctx)
}
