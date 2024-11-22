package iviper

import (
	"github.com/spf13/viper"
	"iconfig"
)

type ViperLoader struct {
}

func (loader *ViperLoader) LoadServerConfig(server *iconfig.ServerConfig) {
	err := viper.UnmarshalKey("server", &server)
	if err != nil {
		panic(err)
	}
}

func (loader *ViperLoader) LoadDatabaseConfig(database *iconfig.DatabaseConfig) {
	err := viper.UnmarshalKey("database", &database)
	if err != nil {
		panic(err)
	}
}

func (loader *ViperLoader) LoadLoggerConfig(logger *iconfig.LoggerConfig) {
	err := viper.UnmarshalKey("logger", &logger)
	if err != nil {
		panic(err)
	}
}

func (loader *ViperLoader) LoadZapConfig(zap *iconfig.ZapConfig) {
	err := viper.UnmarshalKey("logger", &zap.LoggerConfig)
	if err != nil {
		panic(err)
	}
	zap.Compress = viper.GetBool("logger.compress")
}

// Register 注册 loader
// - file：配置文件路径
func Register(file string) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	iconfig.Loader = &ViperLoader{}
}
