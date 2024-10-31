package service

import (
	"github.com/kaibling/apilib/logging"
	"github.com/kaibling/apilib/logging/zap"
)

func BuildLogger(logDriver string) logging.LogWriter {
	switch logDriver {
	case "zap":
		return logging.New(zap.New())
	default:
		return logging.New(zap.New())
	}
}
