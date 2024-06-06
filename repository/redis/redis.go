package redis

import (
	"context"
)

type Config struct {
}

type Repository struct {
	// declare redis client
}

func NewRepository(ctx context.Context, cfg *Config) *Repository {
	//init redis client
	return &Repository{}
}

func (r *Repository) Ping(ctx context.Context) error {
	//tood
	return nil
}
