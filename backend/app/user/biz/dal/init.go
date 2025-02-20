package dal

import (
	"byte_go/backend/app/user/biz/dal/mysql"
	"byte_go/backend/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
