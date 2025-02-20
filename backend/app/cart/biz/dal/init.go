package dal

import (
	"byte_go/backend/app/cart/biz/dal/mysql"
	"byte_go/backend/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
