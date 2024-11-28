package izerolog

import (
	"fmt"
	"github.com/rs/zerolog"
	"iconfig"
	"ilogger"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type zeroLogger struct {
	logger *zerolog.Logger
}

func (logger *zeroLogger) Debug(msg string, args []interface{}) {
	logger.logger.Debug().Msgf(msg, args...)
}

func (logger *zeroLogger) Info(msg string, args []interface{}) {
	logger.logger.Info().Msgf(msg, args...)
}

func (logger *zeroLogger) Warn(msg string, args []interface{}) {
	logger.logger.Warn().Msgf(msg, args...)
}

func (logger *zeroLogger) Error(msg string, args []interface{}) {
	logger.logger.Error().Msgf(msg, args...)
}

func (logger *zeroLogger) Close() {
}

func Register() {
	ilogger.LoggerRegisterFunc = func() (ilogger.Logger, io.Writer) {
		var level zerolog.Level
		if iconfig.Logger.Level == ilogger.DebugLevel {
			level = zerolog.DebugLevel
		} else if iconfig.Logger.Level == ilogger.WarnLevel {
			level = zerolog.WarnLevel
		} else if iconfig.Logger.Level == ilogger.ErrorLevel {
			level = zerolog.ErrorLevel
		} else {
			level = zerolog.InfoLevel
		}

		writer := zeroLogWriter()
		logger := zerolog.New(writer).With().Timestamp().Caller().CallerWithSkipFrameCount(4).Logger()
		logger = logger.Level(level)

		return &zeroLogger{&logger}, writer
	}
}

func zeroLogWriter() zerolog.ConsoleWriter {
	writer := zerolog.ConsoleWriter{}
	writer.NoColor = true
	writer.TimeFormat = time.DateTime
	writer.FormatLevel = func(level interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-5s", level))
	}
	writer.FormatCaller = func(caller interface{}) string {
		var c string
		if cc, ok := caller.(string); ok {
			c = cc
		}
		if len(c) > 0 {
			if cwd, err := os.Getwd(); err == nil {
				if rel, err := filepath.Rel(cwd, c); err == nil {
					c = rel
				}
			}
		}
		return c
	}
	if iconfig.Logger.Out == ilogger.ConsoleLog {
		writer.Out = os.Stdout
	} else {
		writer.Out = ilogger.FileWriter()
	}
	return writer
}
