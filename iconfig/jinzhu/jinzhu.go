package jinzhu

import (
	"github.com/jinzhu/configor"
	"iconfig"
)

type Configuration struct {
	Server struct {
		Mode string `default:"dev"`
		Addr string
		Port string `default:"8080"`
	}

	Database struct {
		Driver      string
		Url         string
		AutoDdl     bool
		MaxIdle     int `default:"5"`
		MaxOpen     int `default:"100"`
		MaxLifeTime int `default:"10"`
	}

	Logger struct {
		Out          string `default:"file"`
		Level        string `default:"info"`
		MaxSize      int    `default:"10"`
		MaxAge       int    `default:"30"`
		MaxBackups   int    `default:"0"`
		Compress     bool   `default:"false"`
		LoggerNumber uint64 `default:"0"`
	}
}

type JinzhuLoader struct {
	config *Configuration
}

func (loader *JinzhuLoader) LoadServerConfig(server *iconfig.ServerConfig) {
	server.Mode = loader.config.Server.Mode
	server.Addr = loader.config.Server.Addr
	server.Port = loader.config.Server.Port
}

func (loader *JinzhuLoader) LoadDatabaseConfig(database *iconfig.DatabaseConfig) {
	database.Driver = loader.config.Database.Driver
	database.Url = loader.config.Database.Url
	database.AutoDdl = loader.config.Database.AutoDdl
	database.MaxIdle = loader.config.Database.MaxIdle
	database.MaxOpen = loader.config.Database.MaxOpen
	database.MaxLifeTime = loader.config.Database.MaxLifeTime
}

func (loader *JinzhuLoader) LoadLoggerConfig(logger *iconfig.LoggerConfig) {
	logger.Out = loader.config.Logger.Out
	logger.Level = loader.config.Logger.Level
	logger.MaxSize = loader.config.Logger.MaxSize
	logger.MaxAge = loader.config.Logger.MaxAge
	logger.MaxBackups = loader.config.Logger.MaxBackups
	logger.LoggerNumber = loader.config.Logger.LoggerNumber
}

func (loader *JinzhuLoader) LoadZapConfig(zap *iconfig.ZapConfig) {
	zap.Out = loader.config.Logger.Out
	zap.Level = loader.config.Logger.Level
	zap.MaxSize = loader.config.Logger.MaxSize
	zap.MaxAge = loader.config.Logger.MaxAge
	zap.MaxBackups = loader.config.Logger.MaxBackups
	zap.LoggerNumber = loader.config.Logger.LoggerNumber
	zap.Compress = loader.config.Logger.Compress
}

// Register 注册 loader
// - file：配置文件路径
func Register(file string) {
	var conf = &Configuration{}
	err := configor.New(&configor.Config{}).Load(conf, file)
	if err != nil {
		panic(err)
	}
	iconfig.Loader = &JinzhuLoader{config: conf}
}
