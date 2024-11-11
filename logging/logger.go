package logging

type LogData struct {
	ReqId          string
	URL            string
	HttpStatusCode int
	Duration       int
	UserName       string
	Method         string
	RequestBody    any
	ResponseBody   any
}

type Writer interface {
	LogRequest(data LogData)
	Info(msg string)
	Error(err error)
	ErrorMsg(msg string)
	Warn(msg string)
	Debug(msg string)
	AddStringField(key string, value string)
	AddIntField(key string, value int)
	AddAnyField(key string, value any)
	NewScope(value string) Writer
}

// type Logger struct {
// 	lw     LogWriter
// 	fields map[string]any
// }

// func New(lw LogWriter) *Logger {
// 	return &Logger{lw: lw, fields: map[string]any{}}
// }

// func (l *Logger) LogRequest(logData LogData) {
// 	l.lw.LogRequest(logData)
// }

// func (l *Logger) Warn(msg string) {
// 	l.lw.Warn(msg)
// }

// func (l *Logger) Info(msg string) {
// 	l.lw.Info(msg)
// }

// func (l *Logger) Error(msg string) {
// 	l.lw.Error(msg)
// }

// func (l *Logger) Debug(msg string) {
// 	l.lw.Debug(msg)
// }
