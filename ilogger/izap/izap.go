package izap

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"iconfig"
	"ilogger"
	"io"
	"os"
	"time"
)

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

func Register() {
	ilogger.LoggerRegisterFunc = func() (ilogger.Logger, io.Writer) {
		var level zapcore.LevelEnabler
		if iconfig.Zap.Level == ilogger.DebugLevel {
			level = zapcore.DebugLevel
		} else if iconfig.Zap.Level == ilogger.WarnLevel {
			level = zapcore.WarnLevel
		} else if iconfig.Zap.Level == ilogger.ErrorLevel {
			level = zapcore.ErrorLevel
		} else {
			level = zapcore.InfoLevel
		}

		writer := zapWriteSyncer()
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

func zapWriteSyncer() zapcore.WriteSyncer {
	if iconfig.Zap.Out == ilogger.ConsoleLog {
		return zapcore.AddSync(os.Stdout)
	} else {
		loggerNumber := iconfig.Zap.LoggerNumber
		if loggerNumber == 0 {
			loggerNumber = ilogger.GetLoggerNumber()
		}

		name := fmt.Sprintf(ilogger.LogDir+"log_%d.log", loggerNumber)

		_logger := &lumberjack.Logger{
			Filename: name,
			// 在进行切割之前，日志文件的最大大小 MB
			MaxSize: iconfig.Zap.MaxSize,
			// 保留旧文件的最大个数
			MaxBackups: iconfig.Zap.MaxBackups,
			// 保留旧文件的最大天数
			MaxAge: iconfig.Zap.MaxAge,
			// 是否压缩/归档旧文件
			Compress: iconfig.Zap.Compress,
			// 使用本地时间
			LocalTime: true,
		}
		return zapcore.AddSync(_logger)
	}
}
