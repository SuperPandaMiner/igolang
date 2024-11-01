package ilogger

import (
	"io"
)

// 日志配置，先注册 LoggerRegisterFunc，后调用 Init() 初始化日志

const (
	ConsoleLog = "console"
	FileLog    = "file"

	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"

	LogDir = "log/"
)

var LoggerRegisterFunc func() (Logger, io.Writer)
var logger Logger
var loggerWriter io.Writer

type Logger interface {
	Debug(msg string, args []interface{})
	Info(msg string, args []interface{})
	Warn(msg string, args []interface{})
	Error(msg string, args []interface{})
	Close()
}

// Init 完成 logger，loggerWriter 的初始化
func Init() {
	if LoggerRegisterFunc == nil {
		panic("LoggerRegisterFunc has not been initialized")
	}
	logger, loggerWriter = LoggerRegisterFunc()
}

func LoggerWriter() io.Writer {
	if loggerWriter == nil {
		panic("LoggerWriter has not been initialized")
	}
	return loggerWriter
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args)
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args)
}

func Close() {
	logger.Close()
}
