package logger

type Logger interface {
	Debug(msg string, data []LogData)
	Info(msg string, data []LogData)
	Warn(msg string, data []LogData)
	Error(msg string, data []LogData)
}

type LogData struct {
	Key string
	Val any
}
