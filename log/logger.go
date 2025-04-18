package log

type Field struct {
	Key   string
	Value any
}

func NewField(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

type LogData struct {
	RequestID      string
	URL            string
	HTTPStatusCode int
	Duration       int
	UserName       string
	Method         string
	RequestBody    any
	ResponseBody   any
}

type Writer interface {
	LogRequest(data LogData)
	New(name string, fields ...Field) Writer
	Named(name string) Writer
	With(fields ...Field) Writer
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Debug(format string, args ...any)
	Error(format string, err error, args ...any)
	Sync()
}
