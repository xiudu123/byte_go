package dal

import (
	"byte_go/backend/app/frontend/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
