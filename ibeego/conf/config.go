package conf

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

type DatabaseConfig struct {
	// 数据库类型
	Dialect string
	// 数据库地址
	Host string
	// 数据库名称
	Dbname string
	// 用户名称
	Username string
	// 密码
	Password string
	// 是否自动执行 ddl
	Autoddl bool
	// 最大空闲数
	MaxIdle int
	// 最大连接数
	MaxOpen int
	// 最大存活，秒
	MaxLifetime int
}

func (conf *DatabaseConfig) load() {
	sectionName := "database::"

	conf.Dialect = String(sectionName + "dialect")
	conf.Host = String(sectionName + "host")
	conf.Dbname = String(sectionName + "dbname")
	conf.Username = String(sectionName + "username")
	conf.Password = String(sectionName + "password")
	conf.Autoddl = Bool(sectionName + "autoddl")
	conf.MaxIdle = config.DefaultInt(sectionName+"maxIdle", 5)
	conf.MaxOpen = config.DefaultInt(sectionName+"maxOpen", 100)
	conf.MaxLifetime = config.DefaultInt(sectionName+"maxLifetime", 10*60)
}

type LoggerConfig struct {
	// 输出，console / file，默认 console
	Out string
	// 日志等级，debug / info / warn / error，默认 info
	Level string
	// 日志文件的最大大小，默认 10 MB
	Maxsize int
	// 文件保存最大天数，默认 7 天
	Maxdays int
}

func (conf *LoggerConfig) load() {
	sectionName := "logger::"

	conf.Out = config.DefaultString(sectionName+"out", logs.AdapterConsole)
	conf.Level = String(sectionName + "level")
	// 单位为 bytes
	conf.Maxsize = config.DefaultInt(sectionName+"maxsize", 10*1024*1024)
	conf.Maxdays = config.DefaultInt(sectionName+"maxdays", 7)
}

var Database = &DatabaseConfig{}
var Logger = &LoggerConfig{}

func Init() {
	Database.load()
	Logger.load()
}

func String(key string) string {
	return config.DefaultString(key, "")
}

func Int(key string) int {
	return config.DefaultInt(key, 0)
}

func Bool(key string) bool {
	return config.DefaultBool(key, false)
}
