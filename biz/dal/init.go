package dal

import (
	"github.com/li-jin-gou/hz_demo/biz/dal/mysql"
	"github.com/li-jin-gou/hz_demo/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
