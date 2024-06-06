package configs

import (
	"gobaseservice/controller"
	"gobaseservice/repository"
	"gobaseservice/service"

	"github.com/gofreego/goutils/logger"
)

type Configuration struct {
	Name       string
	Logger     logger.Config
	Controller controller.Config
	Service    service.Config
	Repository repository.Config
}
