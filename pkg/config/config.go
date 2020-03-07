package config

import (
	"github.com/namsral/flag"
)

type Config struct {
	Usage    bool
	Port     string
	LogLevel string
	Mode     string
}

//GetConfig returns main application config
func GetConfig() Config {
	cfg := Config{}

	flag.StringVar(&cfg.Port, "APP_PORT", "9081", "Port server")

	flag.StringVar(&cfg.LogLevel, "LOG_LEVEL", "debug", "log level")
	flag.StringVar(&cfg.Mode, "MODE", "production", "env app")

	flag.Parse()

	return cfg
}

func Usage() {
	flag.Usage()
}
