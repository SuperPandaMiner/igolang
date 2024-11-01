package iorm

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"ibeego/conf"
	"time"
)

func Init() {
	dialect := conf.Database.Dialect
	dsn := fmt.Sprintf(dialectDSN(dialect), conf.Database.Username, conf.Database.Password,
		conf.Database.Host, conf.Database.Dbname)
	err := orm.RegisterDataBase("default", dialect, dsn,
		orm.MaxIdleConnections(conf.Database.MaxIdle), orm.MaxOpenConnections(conf.Database.MaxOpen), orm.ConnMaxLifetime(time.Duration(conf.Database.MaxLifetime)*time.Second))
	if err != nil {
		panic(err)
	}

	driver = dialect

	if conf.Database.Autoddl {
		if len(AutoMigrateModels) != 0 {
			orm.RegisterModel(AutoMigrateModels...)
			err = orm.RunSyncdb("default", false, true)
			if err != nil {
				panic(err)
			}
		}
	}

	if web.BConfig.RunMode == web.DEV {
		orm.Debug = true
	}
}

func dialectDSN(dialect string) (dsn string) {
	switch dialect {
	case "mysql":
		return "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"
	default:
		panic("unsupported dialect")
	}
}
