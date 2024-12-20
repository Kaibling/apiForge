package zap

import (
	"github.com/kaibling/apiforge/config"
	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func loglevel(level string) zapcore.Level {
	var logLevel zapcore.Level

	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	return logLevel
}

func New(logLevel string, json bool) *Logger {
	var encoding string
	if json {
		encoding = "json"
	} else {
		encoding = "console"
	}

	cfg := zap.Config{
		Encoding:      encoding,                                 // Output format: json or console
		Level:         zap.NewAtomicLevelAt(loglevel(logLevel)), // Set the default log level
		OutputPaths:   []string{"stdout"},                       // Write logs to stdout
		EncoderConfig: zap.NewProductionEncoderConfig(),         // Encoder configuration
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		// todo remove panic
		panic(err)
	}

	return &Logger{
		ID:       utils.NewULID().String(),
		l:        logger,
		Fields:   map[string]zapcore.Field{},
		logLevel: logLevel,
		json:     json,
	}
}

func logCopy(log *Logger) *Logger {
	newLog := New(log.logLevel, log.json)

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
	json     bool
}

func (l *Logger) LogRequest(data logging.LogData) {
	requestFields := []zapcore.Field{
		zap.String("url", data.URL),
		zap.Int("duration_ms", data.Duration),
		zap.Int("http_status_code", data.HTTPStatusCode),
		zap.String("method", data.Method),
	}

	if config.LogRequestBody {
		requestFields = append(requestFields, zap.Any("request_body", data.RequestBody))
	}

	if config.LogResponseBody {
		requestFields = append(requestFields, zap.Any("response_body", data.ResponseBody))
	}

	requestFields = append(requestFields, l.fields()...)
	l.l.Info("request", requestFields...)
}

func (l *Logger) fields() []zapcore.Field {
	values := []zapcore.Field{}
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

func (l *Logger) NewScope(value string) logging.Writer { //nolint: ireturn, nolintlint
	newLogger := logCopy(l)
	newLogger.Fields["scope"] = zap.Any("scope", value)

	return newLogger
}
