package repository

import (
	"context"
	"gobaseservice/internal/repository/redis"
	"gobaseservice/internal/service"
)

const (
	REDIS = "redis"
	MONGO = "mongo"
	MYSQL = "mysql"
)

type Config struct {
	Name  string
	Redis redis.Config
}

func NewRepository(ctx context.Context, cfg *Config) service.Repository {
	switch cfg.Name {
	case REDIS:
		return redis.NewRepository(ctx, &cfg.Redis)
	}
	panic("invalid repository name , provided `" + cfg.Name + "` expected + `" + REDIS + "`")
}
