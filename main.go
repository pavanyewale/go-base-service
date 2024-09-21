package main

import (
	"context"
	"flag"
	"fmt"
	"gobaseservice/cmd/grpc_server"
	"gobaseservice/cmd/http_server"
	"gobaseservice/internal/configs"
	"gobaseservice/internal/constants"

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
	Run(ctx context.Context)
	apputils.Application
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
	var apps []Application
	for _, appName := range conf.AppNames {
		switch appName {
		case constants.HTTP_SERVER:
			apps = append(apps, http_server.NewHTTPServer(&conf))
		case constants.GRPC_SERVER:
			apps = append(apps, grpc_server.NewGRPCServer(&conf))
		default:
			logger.Panic(ctx, "invalid application name provided `%s`", appName)
		}
	}
	apps_to_graceful_shutdown := make([]apputils.Application, len(apps))
	for _, app := range apps {
		logger.Info(ctx, "Starting %s", app.Name())
		go app.Run(ctx)
		apps_to_graceful_shutdown = append(apps_to_graceful_shutdown, app)
	}

	apputils.GracefulShutdown(ctx, apps_to_graceful_shutdown...)
}
