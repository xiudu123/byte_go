package main

import (
	"byte_go/backend/utils"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/**
 * @author: 锈渎
 * @date: 2025/2/23 21:46
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 测试
 */
var rdb *redis.Client

func ini() {
	// 创建一个 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
		DB:   1,                // 使用默认数据库
	})
}

func main() {
	ini()
	mist := utils.NewMist(rdb, "test")
	for i := 0; i < 10; i++ {
		id, err := mist.GenerateId()
		fmt.Println(id)
		fmt.Println(err)
	}

}
