package redis

import (
	"byte_go/backend/app/auth/conf"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func SetJTI(ctx context.Context, userId uint32, jti string) (err error) {
	return RedisClient.SAdd(ctx, "jti_list:"+strconv.Itoa(int(userId)), jti).Err()
}

func ListJTIList(ctx context.Context, userId uint32) ([]string, error) {
	return RedisClient.SMembers(ctx, "jti_list:"+strconv.Itoa(int(userId))).Result()
}

func AddJTI2BlackListed(ctx context.Context, jti string) (err error) {
	return RedisClient.SAdd(ctx, "blacklisted:jti", jti).Err()
}

func AddJTIList2BlackListed(ctx context.Context, jtiList []string) (err error) {
	return RedisClient.SAdd(ctx, "blacklisted:jti", jtiList).Err()
}

func CheckJITIsBlackListed(ctx context.Context, jti string) (bool, error) {
	return RedisClient.SIsMember(ctx, "blacklisted:jti", jti).Result()
}
