module github.com/hertz/hello

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/apache/thrift v0.0.0-00010101000000-000000000000
	github.com/cloudwego/hertz v0.4.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/hertz-contrib/gzip v0.0.1
	github.com/hertz-contrib/logger/accesslog v0.0.0-20221206155315-4315a5ae6b90
	github.com/hertz-contrib/logger/logrus v1.0.0
	github.com/hertz-contrib/pprof v0.1.0
	github.com/kr/pretty v0.3.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)
