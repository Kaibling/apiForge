package zap

import (
	"github.com/kaibling/apiforge/logging"
	"go.uber.org/zap"
)

func New() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		l: logger,
	}
	//defer logger.Sync()

}

type Logger struct {
	l *zap.Logger
}

func (l *Logger) LogRequest(data logging.LogData) {
	l.l.Info("request",
		zap.String("req_id", data.ReqId),
		zap.String("url", data.URL),
		zap.String("user", data.UserName),
		zap.Int("duration", data.Duration),
		zap.Int("http_status_code", data.HttpStatusCode),
		zap.Any("request_body", data.RequestBody),
		zap.Any("response_body", data.ResponseBody),
		zap.String("method", data.Method),
	)
}

func (l *Logger) LogLine(msg string) {
	l.l.Info(msg)
}
