package dal

import (
	"byte_go/backend/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
