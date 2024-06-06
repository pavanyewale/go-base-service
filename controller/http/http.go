package http

import (
	"context"
	"fmt"
	"gobaseservice/controller/http/ping"
	"gobaseservice/controller/http/swagger"
	"gobaseservice/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofreego/goutils"
	"github.com/gofreego/goutils/logger"
)

type Config struct {
	Port               int
	GinMode            string
	ReadTimeoutMillis  int
	WriteTimeoutMillis int
	Ping               ping.Config
}

type Controller struct {
	server         *http.Server
	conf           *Config
	serviceFactory *service.ServiceFactory
}

func NewController(ctx context.Context, c *Config, sf *service.ServiceFactory) *Controller {
	return &Controller{conf: c, serviceFactory: sf}
}
func (c *Controller) Name() string {
	return "HTTP controller"
}

func (c *Controller) registerHandlers(ctx context.Context, router gin.IRouter) {
	ping.NewHandler(ctx, &c.conf.Ping, c.serviceFactory.PingService).Register(router)
	swagger.NewHandler(ctx).Register(router)
}

func (c *Controller) Listen(ctx context.Context) error {
	router := goutils.GetHTTPRouter(c.conf.GinMode)

	c.registerHandlers(ctx, router.Group("/gobaseservice"))

	logger.Info(ctx, "🌏 go-base-service started on 🌎 -> http://localhost:%d/", c.conf.Port)

	c.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.conf.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(c.conf.ReadTimeoutMillis) * time.Millisecond,
		WriteTimeout: time.Duration(c.conf.WriteTimeoutMillis) * time.Millisecond,
	}
	err := c.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Panic(ctx, "Failed to start server : %v", err)
		return err
	}
	return nil
}

func (c *Controller) Shutdown(ctx context.Context) {
	err := c.server.Shutdown(ctx)
	if err != nil {
		logger.Error(ctx, "Failed to shutdown http server : %v", err)
		return
	}
}
