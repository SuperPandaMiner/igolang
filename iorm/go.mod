module iorm

go 1.22

require (
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	iconfig v0.0.0
	ilogger v0.0.0
	utils v0.0.0
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/configor v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace iconfig => ../iconfig

replace ilogger => ../ilogger

replace utils => ../utils
