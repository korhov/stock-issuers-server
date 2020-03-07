package main

import (
	"github.com/korhov/stock-issuers-server/configs"
	internalLog "github.com/korhov/stock-issuers-server/internal/log"
	internalRoutes "github.com/korhov/stock-issuers-server/internal/routes"
	"github.com/korhov/stock-issuers-server/internal/tracing"

	"net/http"
)

func main() {
	appCfg := configs.GetConfig()

	if appCfg.Usage {
		configs.Usage()
		return
	}

	logger, err := internalLog.GetLogger(internalLog.Config{
		LogLevel:      appCfg.LogLevel,
		IsDevelopment: appCfg.Mode != internalLog.ModeProduction,
	})
	if err != nil {
		panic(err)
	}

	tracer, err := tracing.Init(logger, appCfg.Jaeger)

	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":"+appCfg.Port, internalRoutes.Routes(logger, tracer)); err != nil {
		panic(err)
	}
}
