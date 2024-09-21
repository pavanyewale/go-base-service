package configs

import (
	"gobaseservice/internal/repository"
	"gobaseservice/internal/service"

	"github.com/gofreego/goutils/logger"
)

type Configuration struct {
	Name       string
	AppNames   []string
	Logger     logger.Config
	Service    service.Config
	Repository repository.Config
	Server     Server
}

type Server struct {
	GRPCPort int
	HTTPPort int
}
