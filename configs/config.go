package configs

import (
	"github.com/korhov/stock-issuers-server/internal/tracing"
	"github.com/namsral/flag"
)

type Config struct {
	Usage    bool
	Port     string
	LogLevel string
	Mode     string
	Jaeger   tracing.Config
}

//GetConfig returns main application config
func GetConfig() Config {
	cfg := Config{}

	flag.StringVar(&cfg.Port, "APP_PORT", "9081", "Port server")

	flag.StringVar(&cfg.LogLevel, "LOG_LEVEL", "debug", "log level")
	flag.StringVar(&cfg.Mode, "MODE", "production", "env app")

	flag.BoolVar(&cfg.Jaeger.Enabled, "JAEGER_ENABLED", false, "")
	flag.StringVar(&cfg.Jaeger.Reporter, "JAEGER_REPORTER", "", "")
	flag.StringVar(&cfg.Jaeger.ServiceName, "JAEGER_SERVICE_NAME", "", "")

	flag.Parse()

	return cfg
}

func Usage() {
	flag.Usage()
}
