package controller

import (
	"context"
	"gobaseservice/controller/http"
	"gobaseservice/service"
)

type Config struct {
	Name string
	HTTP http.Config
}

type Controller interface {
	Listen(ctx context.Context) error
	Shutdown(ctx context.Context)
	Name() string
}

func NewController(ctx context.Context, c *Config, sf *service.ServiceFactory) Controller {
	switch c.Name {
	case "http":
		return http.NewController(ctx, &c.HTTP, sf)
	}
	panic("unknown controller name provided `" + c.Name + "` expected `http`")
}
