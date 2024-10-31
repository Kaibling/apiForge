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

type LogWriter interface {
	LogRequest(data LogData)
	LogLine(msg string)
}

type Logger struct {
	lw LogWriter
}

func New(lw LogWriter) *Logger {
	return &Logger{lw: lw}
}

func (l *Logger) LogRequest(logData LogData) {
	l.lw.LogRequest(logData)
}

func (l *Logger) LogLine(msg string) {
	l.lw.LogLine(msg)
}
