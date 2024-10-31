package service

import (
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/logging/zap"
)

func BuildLogger(logDriver string) logging.LogWriter {
	switch logDriver {
	case "zap":
		return logging.New(zap.New())
	default:
		return logging.New(zap.New())
	}
}
