package service

import (
	"context"
	"gobaseservice/internal/service/ping"
)

type Config struct {
	Ping ping.Config
}

type Repository interface {
	ping.Repository
}

type ServiceFactory struct {
	PingService *ping.Service
}

func NewServiceFactory(ctx context.Context, cfg *Config, repo Repository) *ServiceFactory {
	sf := ServiceFactory{}
	sf.PingService = ping.NewService(ctx, &cfg.Ping, repo)
	return &sf
}
