package dal

import (
	"byte_go/backend/app/order/biz/dal/mysql"
	"byte_go/backend/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
