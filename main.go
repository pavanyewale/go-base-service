package main

import (
	"context"
	"flag"
	"fmt"
	"gobaseservice/configs"
	"gobaseservice/controller"
	"gobaseservice/repository"
	"gobaseservice/service"

	"github.com/gofreego/goutils/apputils"
	"github.com/gofreego/goutils/configutils"
	"github.com/gofreego/goutils/logger"
	"gopkg.in/yaml.v3"
)

var (
	env  string
	path string
)

type Application interface {
	GetName() string
	Shutdown(ctx context.Context)
}

func main() {
	flag.StringVar(&env, "env", "dev", "-env=dev")
	flag.StringVar(&path, "path", ".", "-path=./")
	flag.Parse()
	ctx := context.Background()

	configfile := fmt.Sprintf("%s/%s.yaml", path, env)

	var conf configs.Configuration
	// bytes, _ := yaml.Marshal(conf)
	// print(string(bytes))
	err := configutils.ReadConfig(ctx, configfile, &conf)
	if err != nil {
		logger.Panic(ctx, "failed to read configs : %v", err)
	}
	// initiating logger
	if err := conf.Logger.InitiateLogger(); err != nil {
		logger.Panic(ctx, "failed to initiate logger, %v", err)
	}
	// logging config for debug
	bytes, _ := yaml.Marshal(conf)
	logger.Debug(ctx, "\n%s", bytes)
	// starting application

	repo := repository.NewRepository(ctx, &conf.Repository)
	serviceFactory := service.NewServiceFactory(ctx, &conf.Service, repo)
	cntlr := controller.NewController(ctx, &conf.Controller, serviceFactory)
	go cntlr.Listen(ctx)
	apputils.GracefulShutdown(ctx, cntlr)
}
