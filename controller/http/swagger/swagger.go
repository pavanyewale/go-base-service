package swagger

import (
	"context"
	_ "gobaseservice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler(ctx context.Context) *Handler {
	return &Handler{}
}

func (h *Handler) Register(router gin.IRouter) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
