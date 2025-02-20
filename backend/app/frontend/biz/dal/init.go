package dal

import (
	"byte_go/backend/app/front/biz/dal/mysql"
	"byte_go/backend/app/front/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
