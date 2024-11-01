package jinzhu

import (
	"github.com/jinzhu/configor"
	"igin/config/iconfig"
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

func (loader *JinzhuLoader) LoadServerConfig() *iconfig.ServerConfig {
	return &iconfig.ServerConfig{
		Mode: loader.config.Server.Mode,
		Addr: loader.config.Server.Addr,
		Port: loader.config.Server.Port,
	}
}

func (loader *JinzhuLoader) LoadDatabaseConfig() *iconfig.DatabaseConfig {
	return &iconfig.DatabaseConfig{
		Driver:      loader.config.Database.Driver,
		Url:         loader.config.Database.Url,
		AutoDdl:     loader.config.Database.AutoDdl,
		MaxIdle:     loader.config.Database.MaxIdle,
		MaxOpen:     loader.config.Database.MaxOpen,
		MaxLifeTime: loader.config.Database.MaxLifeTime,
	}
}

func (loader *JinzhuLoader) LoadLoggerConfig() *iconfig.LoggerConfig {
	return &iconfig.LoggerConfig{
		Out:          loader.config.Logger.Out,
		Level:        loader.config.Logger.Level,
		MaxSize:      loader.config.Logger.MaxSize,
		MaxAge:       loader.config.Logger.MaxAge,
		MaxBackups:   loader.config.Logger.MaxBackups,
		Compress:     loader.config.Logger.Compress,
		LoggerNumber: loader.config.Logger.LoggerNumber,
	}
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
