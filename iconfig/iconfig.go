package iconfig

type ServerConfig struct {
	// gin mode，dev / prod，默认 dev
	Mode string
	// 应用地址
	Addr string
	// 应用端口
	Port string
}

type DatabaseConfig struct {
	// 数据库类型
	Driver string
	// 数据库 url
	Url string
	// 是否自动执行 ddl，默认 false
	AutoDdl bool
	// 空闲连接数，默认 5
	MaxIdle int
	// 最大连接数， 默认 100
	MaxOpen int
	// 连接最大存活时间， 默认 10 分钟
	MaxLifeTime int
}

type LoggerConfig struct {
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
	// 指定日志编号，指定编号则不会读取日志编号文件中的数值
	LoggerNumber uint64
	// 是否开启日志压缩，默认不开启
	Compress bool
}

type ConfigLoader interface {
	LoadServerConfig(*ServerConfig)
	LoadDatabaseConfig(*DatabaseConfig)
	LoadLoggerConfig(*LoggerConfig)
}

var Server = &ServerConfig{
	Mode: "dev",
}
var Database = &DatabaseConfig{
	MaxIdle:     5,
	MaxOpen:     100,
	MaxLifeTime: 10,
}
var Logger = &LoggerConfig{
	Out:          "console",
	Level:        "info",
	MaxSize:      10,
	MaxAge:       7,
	MaxBackups:   0,
	LoggerNumber: 0,
}

var Loader ConfigLoader

func Init() {
	if Loader == nil {
		panic("Loader has not been initialized")
	}
	Loader.LoadServerConfig(Server)
	Loader.LoadDatabaseConfig(Database)
	Loader.LoadLoggerConfig(Logger)
}
