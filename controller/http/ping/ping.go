package ping

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gofreego/goutils/response"
)

type Config struct {
}

type Service interface {
	Ping(ctx context.Context) error
}
type Handler struct {
	cfg     *Config
	service Service
}

func NewHandler(ctx context.Context, cfg *Config, service Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) Register(router gin.IRouter) {
	router.GET("/ping", h.Ping)
}

// Ping godoc
// @Summary Ping the server
// @Description Check if the server is alive
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} string "Okay, I am alive!"
// @Failure 500 {object} error "Internal Server Error"
// @Router /gobaseservice/ping [get]
func (h *Handler) Ping(ctx *gin.Context) {
	err := h.service.Ping(ctx)
	if err != nil {
		response.WriteError(ctx, err)
		return
	}
	response.WriteSuccess(ctx, "Okay, I am alive!")
}
