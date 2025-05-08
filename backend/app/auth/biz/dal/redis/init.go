package redis

import (
	"byte_go/backend/app/auth/conf"
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

type defaultHook struct{}

func (d *defaultHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}
func (d *defaultHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		err := next(ctx, cmd)
		if err == redis.Nil {
			switch v := cmd.(type) {
			case *redis.StringCmd:
				v.SetVal("0")
			case *redis.IntCmd:
				v.SetVal(0)
			}
		}
		return nil
	}
}

func (d *defaultHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return next
}

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})

	// 添加默认值处理包装器
	RedisClient.AddHook(&defaultHook{})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

}
