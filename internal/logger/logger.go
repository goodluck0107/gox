package logger

type Logger interface {
	StartUp(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}
