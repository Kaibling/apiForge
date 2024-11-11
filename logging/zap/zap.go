package zap

import (
	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func loglevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}

func New(logLevel string) *Logger {
	cfg := zap.Config{
		Encoding:      "json",                                   // Output format: json or console
		Level:         zap.NewAtomicLevelAt(loglevel(logLevel)), // Set the default log level
		OutputPaths:   []string{"stdout"},                       // Write logs to stdout
		EncoderConfig: zap.NewProductionEncoderConfig(),         // Encoder configuration
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		ID:       utils.NewULID().String(),
		l:        logger,
		Fields:   map[string]zapcore.Field{},
		logLevel: logLevel,
	}
	//defer logger.Sync()

}

func logCopy(log *Logger) *Logger {
	newLog := New(log.logLevel)
	//newLog.Fields = make(map[string]zapcore.Field)
	for k, v := range log.Fields {
		newLog.Fields[k] = v
	}

	return newLog
}

type Logger struct {
	ID       string
	l        *zap.Logger
	logLevel string
	Fields   map[string]zapcore.Field
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

func (l *Logger) fields() []zapcore.Field {
	var values []zapcore.Field
	for _, value := range l.Fields {
		values = append(values, value)
	}

	return values
}

func (l *Logger) AddStringField(key string, value string) {
	l.Fields[key] = zap.String(key, value)
}
func (l *Logger) AddIntField(key string, value int) {
	l.Fields[key] = zap.Int(key, value)
}
func (l *Logger) AddAnyField(key string, value any) {
	l.Fields[key] = zap.Any(key, value)
}

func (l *Logger) ErrorMsg(msg string) {
	l.l.Error(msg, l.fields()...)
}

func (l *Logger) Error(err error) {
	l.l.Error(err.Error(), l.fields()...)
}

func (l *Logger) Debug(msg string) {
	l.l.Debug(msg, l.fields()...)
}

func (l *Logger) Warn(msg string) {
	l.l.Warn(msg, l.fields()...)
}

func (l *Logger) Info(msg string) {
	l.l.Info(msg, l.fields()...)
}

func (l *Logger) NewScope(value string) logging.Writer {
	newLogger := logCopy(l)
	newLogger.Fields["scope"] = zap.Any("scope", value)

	return newLogger
}
