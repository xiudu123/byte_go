package dal

import (
	"byte_go/backend/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	//mysql.Init()
}
