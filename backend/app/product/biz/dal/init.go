package dal

import (
	"byte_go/backend/app/product/biz/dal/mysql"
	"byte_go/backend/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
