package ilogger

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"iconfig"
	"io"
	"os"
	"strconv"
)

// 日志配置，先注册 LoggerRegisterFunc，后调用 Init() 初始化日志

const (
	ConsoleLog = "console"
	FileLog    = "file"

	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"

	logDir = "log/"
)

const loggerNumberFile = logDir + "loggerNumber.txt"

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

func FileWriter() io.Writer {
	loggerNumber := iconfig.Logger.LoggerNumber
	if loggerNumber == 0 {
		loggerNumber = getLoggerNumber()
	}

	name := fmt.Sprintf(logDir+"log_%d.log", loggerNumber)

	writer := &lumberjack.Logger{
		Filename: name,
		// 在进行切割之前，日志文件的最大大小 MB
		MaxSize: iconfig.Logger.MaxSize,
		// 保留旧文件的最大个数
		MaxBackups: iconfig.Logger.MaxBackups,
		// 保留旧文件的最大天数
		MaxAge: iconfig.Logger.MaxAge,
		// 是否压缩/归档旧文件
		Compress: iconfig.Logger.Compress,
		// 使用本地时间
		LocalTime: true,
	}
	return writer
}

// getLoggerNumber 扫描文件获取本次的日志编号
func getLoggerNumber() uint64 {
	var number uint64
	fileExist := true
	file, err := os.OpenFile(loggerNumberFile, os.O_RDWR, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fileExist = false
		} else {
			panic(err)
		}
	}
	defer file.Close()

	if !fileExist {
		number = 0

		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			panic(err)
		}
		file, err = os.Create(loggerNumberFile)
		if err != nil {
			panic(err)
		}
	} else {
		scanner := bufio.NewScanner(file)
		// 扫描第一行
		if scanner.Scan() {
			number, err = strconv.ParseUint(scanner.Text(), 10, 0)
			if err != nil {
				number = 0
			} else {
				number += 1
			}
		} else {
			number = 0
		}
	}

	// 清空文件
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// 重置指针
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// 写入编号
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%d\n", number))
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return number
}
