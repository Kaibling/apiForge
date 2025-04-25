package zap

import (
	"fmt"

	"github.com/kaibling/apiforge/config"
	"github.com/kaibling/apiforge/log"
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

type LogConfig struct {
	JSONLogging bool
	LogLevel    string // Log level: "debug", "info", "warn", "error"
	AppName     string
}

// Init initializes a new AppLogger with the given configuration and optional fields.
func New(config LogConfig, fields ...log.Field) *Logger {
	baseLogger := initLogger(config).Named(config.AppName)
	zapFields := unwrapFields(fields)
	return &Logger{
		l: baseLogger.With(zapFields...).WithOptions(zap.AddCallerSkip(1)),
	}
}

func initLogger(cfg LogConfig) *zap.Logger {
	var encoding string
	if cfg.JSONLogging {
		encoding = "json"
	} else {
		encoding = "console"
	}

	logcfg := zap.Config{
		Encoding:      encoding,                                     // Output format: json or console
		Level:         zap.NewAtomicLevelAt(loglevel(cfg.LogLevel)), // Set the default log level
		OutputPaths:   []string{"stdout"},                           // Write logs to stdout
		EncoderConfig: zap.NewProductionEncoderConfig(),             // Encoder configuration
	}
	logcfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := logcfg.Build()
	if err != nil {
		// todo remove panic
		panic(err)
	}

	return logger
}

type Logger struct {
	l      *zap.Logger
	Fields map[string]zapcore.Field
}

func (l *Logger) LogRequest(data log.LogData) {
	requestFields := []log.Field{
		{Key: "url", Value: data.URL},
		{Key: "duration_ms", Value: data.Duration},
		{Key: "http_status_code", Value: data.HTTPStatusCode},
		{Key: "method", Value: data.Method},
	}

	if config.LogRequestBody {
		requestFields = append(requestFields, log.NewField("request_body", data.RequestBody))
	}

	if config.LogResponseBody {
		requestFields = append(requestFields, log.Field{Key: "response_body", Value: data.ResponseBody})
	}

	l.With(requestFields...).Info("request")
}

// New creates a new AppLogger with additional fields and a name.
func (cl *Logger) New(name string, fields ...log.Field) log.Writer {
	return cl.With(fields...).Named(name)
}

// With adds fields to the logger and returns a new AppLogger.
func (cl *Logger) With(fields ...log.Field) log.Writer {
	zapFields := unwrapFields(fields)
	return &Logger{
		l: cl.l.With(zapFields...),
	}
}

// Named creates a new AppLogger with a specific name.
func (cl *Logger) Named(name string) log.Writer {
	return &Logger{
		l: cl.l.Named(name),
	}
}

// Infof logs an informational message with formatting.
func (cl *Logger) Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	cl.l.Info(msg)
}

// Warnf logs a warning message with formatting.
func (cl *Logger) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	cl.l.Warn(msg)
}

// Debugf logs a debug message with formatting.
func (cl *Logger) Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	cl.l.Debug(msg)
}

// Errorf logs an error message with formatting.
func (cl *Logger) Error(format string, err error, args ...any) {
	msg := fmt.Sprintf(format, args...)
	cl.l.Error(msg, zap.Error(err))
}

// Sync flushes any buffered log entries.
func (cl *Logger) Sync() {
	cl.l.Sync()
}

// unwrapFields converts a slice of Field to a slice of zap.Field.
func unwrapFields(fields []log.Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}
