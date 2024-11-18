module ilogger

go 1.22

require (
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.27.0
	iconfig v0.0.0
)

require (
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace iconfig => ../iconfig
