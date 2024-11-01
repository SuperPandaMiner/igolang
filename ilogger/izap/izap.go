package izap

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ilogger"
	"io"
	"os"
	"strconv"
	"time"
)

const loggerNumberFile = ilogger.LogDir + "loggerNumber.txt"

type zapLogger struct {
	logger *zap.SugaredLogger
}

func (logger *zapLogger) Debug(msg string, args []interface{}) {
	logger.logger.Debugf(msg, args...)
}

func (logger *zapLogger) Info(msg string, args []interface{}) {
	logger.logger.Infof(msg, args...)
}

func (logger *zapLogger) Warn(msg string, args []interface{}) {
	logger.logger.Warnf(msg, args...)
}

func (logger *zapLogger) Error(msg string, args []interface{}) {
	logger.logger.Errorf(msg, args...)
}

func (logger *zapLogger) Close() {
	logger.logger.Sync()
}

type Config struct {
	// 输出，console / file，默认 console
	Out string
	// 日志等级，debug / info / warn / error，默认 info
	Level string
	// 日志文件的最大大小，默认 10 MB
	MaxSize int
	// 文件保存最大天数，默认 7 天
	MaxAge int
	// 保留旧文件的最大个数，默认不限制
	MaxBackups int
	// 是否开启日志压缩，默认不开启
	Compress bool
	// 指定日志编号，指定编号则不会读取日志编号文件中的数值
	LoggerNumber uint64
}

func Register(config *Config) {
	ilogger.LoggerRegisterFunc = func() (ilogger.Logger, io.Writer) {
		var level zapcore.LevelEnabler
		if config.Level == ilogger.DebugLevel {
			level = zapcore.DebugLevel
		} else if config.Level == ilogger.WarnLevel {
			level = zapcore.WarnLevel
		} else if config.Level == ilogger.ErrorLevel {
			level = zapcore.ErrorLevel
		} else {
			level = zapcore.InfoLevel
		}

		writer := zapWriteSyncer(config)
		core := zapcore.NewCore(zapEncoder(), writer, level)

		// AddCaller() 启用 caller 配置；AddCallerSkip() 跳过包装的最后两层路径，显示调用路径
		_logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
		sugaredLogger := _logger.Sugar()

		return &zapLogger{sugaredLogger}, writer
	}
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "Message",
		StacktraceKey:  "StackTrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(config *Config) zapcore.WriteSyncer {
	if config.Out == ilogger.ConsoleLog {
		return zapcore.AddSync(os.Stdout)
	} else {
		loggerNumber := config.LoggerNumber
		if loggerNumber == 0 {
			loggerNumber = getLoggerNumber()
		}
		if config.MaxSize == 0 {
			config.MaxSize = 10
		}
		if config.MaxAge == 0 {
			config.MaxAge = 7
		}

		name := fmt.Sprintf(ilogger.LogDir+"log_%d.log", loggerNumber)

		_logger := &lumberjack.Logger{
			Filename: name,
			// 在进行切割之前，日志文件的最大大小 MB
			MaxSize: config.MaxSize,
			// 保留旧文件的最大个数
			MaxBackups: config.MaxBackups,
			// 保留旧文件的最大天数
			MaxAge: config.MaxAge,
			// 是否压缩/归档旧文件
			Compress: config.Compress,
			// 使用本地时间
			LocalTime: true,
		}
		return zapcore.AddSync(_logger)
	}
}

// 扫描文件获取本次的日志编号
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

		err = os.MkdirAll(ilogger.LogDir, 0755)
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
