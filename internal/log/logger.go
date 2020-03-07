package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ModeProduction  string = "production"
	ModeDevelopment string = "development"
)

// Config is structure described config of logger
type Config struct {
	IsDevelopment bool
	LogLevel      string
}

func zapLevelFromString(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	}

	return zapcore.InfoLevel
}

func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetLogger creates if need and return zap logger
func GetLogger(cfg Config) (*zap.SugaredLogger, error) {
	logCfg := zap.NewProductionConfig()
	if cfg.IsDevelopment {
		logCfg = zap.NewDevelopmentConfig()
	}
	logCfg.Encoding = "json"
	logCfg.EncoderConfig = newProductionEncoderConfig()
	logCfg.Level.SetLevel(zapLevelFromString(cfg.LogLevel))
	logCfg.Development = cfg.IsDevelopment
	lg, err := logCfg.Build()
	if err != nil {
		return nil, err
	}

	return lg.Sugar(), nil
}
