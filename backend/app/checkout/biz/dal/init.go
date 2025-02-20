package dal

import (
	"byte_go/backend/app/checkout/biz/dal/mysql"
	"byte_go/backend/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
