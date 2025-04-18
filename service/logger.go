package service

import (
	"github.com/kaibling/apiforge/config"
	"github.com/kaibling/apiforge/log"
	"github.com/kaibling/apiforge/log/zap"
)

type LogConfig struct {
	LogDriver    string
	LogLevel     string
	RequestBody  bool
	ResponseBody bool
	JSON         bool
	AppName      string
}

func BuildLogger(cfg LogConfig) log.Writer { //nolint: ireturn, nolintlint
	config.LogRequestBody = cfg.RequestBody
	config.LogResponseBody = cfg.ResponseBody

	switch cfg.LogDriver {
	case "zap":
		return zap.New(zap.LogConfig{
			JSONLogging: cfg.JSON,
			LogLevel:    cfg.LogLevel,
			AppName:     cfg.AppName,
		})
	default:
		return zap.New(zap.LogConfig{
			JSONLogging: cfg.JSON,
			LogLevel:    cfg.LogLevel,
			AppName:     cfg.AppName,
		})
	}
}
