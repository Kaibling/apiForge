package service

import (
	"github.com/kaibling/apiforge/config"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/logging/zap"
)

type LogConfig struct {
	LogDriver    string
	LogLevel     string
	RequestBody  bool
	ResponseBody bool
	JSON         bool
}

func BuildLogger(cfg LogConfig) logging.Writer { //nolint: ireturn, nolintlint
	config.LogRequestBody = cfg.RequestBody
	config.LogResponseBody = cfg.ResponseBody

	switch cfg.LogDriver {
	case "zap":
		return zap.New(cfg.LogLevel, cfg.JSON)
	default:
		return zap.New(cfg.LogLevel, cfg.JSON)
	}
}
