package iorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"igin/config/iconfig"
	"igin/engine"
	"ilogger"
	"log"
	"time"
)

const mysqlDriver = "mysql"

var Close func()

func Init() {
	dialector := dbDialector(iconfig.Database.Driver)
	db := dbConnection(dialector)

	if iconfig.Database.AutoDdl {
		err := db.AutoMigrate(AutoMigrateModels...)
		if err != nil {
			panic(err)
		}
	}

	// igorm db
	gdb = db

	Close = func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func dbDialector(driver string) gorm.Dialector {
	var dialector gorm.Dialector
	switch driver {
	case mysqlDriver:
		config := mysql.Config{
			DSN:               iconfig.Database.Url,
			DefaultStringSize: 255,
		}
		dialector = mysql.New(config)
	default:
		panic("unsupported driver")
	}
	return dialector
}

func dbConnection(dialector gorm.Dialector) *gorm.DB {
	loggerLevel := logger.Info
	if engine.IsModeProd() {
		loggerLevel = logger.Warn
	}
	gormConfig := gorm.Config{
		Logger: logger.New(
			log.New(ilogger.LoggerWriter(), "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  loggerLevel, // 日志级别
				IgnoreRecordNotFoundError: true,        // 日志中忽略 ErrRecordNotFound 错误
				Colorful:                  false,       // 禁用彩色打印
				ParameterizedQueries:      true,        // sql 只打印占位符
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键
	}

	gormdb, err := gorm.Open(
		dialector, &gormConfig,
	)
	if err != nil {
		panic("failed to connect to the database")
	}

	sqldb, err := gormdb.DB()
	if err != nil {
		panic("failed to get the sql db")
	}
	sqldb.SetMaxIdleConns(iconfig.Database.MaxIdle)
	sqldb.SetMaxOpenConns(iconfig.Database.MaxOpen)
	sqldb.SetConnMaxLifetime(time.Duration(iconfig.Database.MaxLifeTime) * time.Minute)

	return gormdb
}
