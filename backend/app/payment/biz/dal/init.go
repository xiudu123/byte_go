package dal

import (
	"byte_go/backend/app/payment/biz/dal/mysql"
	"byte_go/backend/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
