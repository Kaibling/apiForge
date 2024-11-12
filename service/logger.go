package service

import (
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/logging/zap"
)

type LogConfig struct {
	LogDriver string
	LogLevel  string
}

func BuildLogger(cfg LogConfig) logging.Writer { //nolint: ireturn, nolintlint
	switch cfg.LogDriver {
	case "zap":
		return zap.New(cfg.LogLevel)
	default:
		return zap.New(cfg.LogLevel)
	}
}
