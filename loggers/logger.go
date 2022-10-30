package loggers

type Logger interface {
	Debug(params ...interface{}) *Logger
	Info(params ...interface{}) *Logger
	Success(params ...interface{}) *Logger
	Warn(params ...interface{}) *Logger
	Error(params ...interface{}) *Logger
	Fatal(params ...interface{}) *Logger
}
