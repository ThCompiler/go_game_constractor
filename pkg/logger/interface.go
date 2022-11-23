package logger

type LogLevel string

const (
	ErrorLevel LogLevel = "error"
	WarnLevel  LogLevel = "warn"
	InfoLevel  LogLevel = "info"
	DebugLevel LogLevel = "debug"
	PanicLevel LogLevel = "panic"
	FatalLevel LogLevel = "fatal"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message interface{}, args ...interface{})
	Warn(message interface{}, args ...interface{})
	Error(message interface{}, args ...interface{})
	Panic(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
	With(key Field, value interface{}) Interface
}
