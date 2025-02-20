package redis

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"

	"byte_go/backend/app/auth/conf"
	"github.com/redis/go-redis/v9"
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

func SetUserToken(ctx context.Context, userId uint32, token string, ttl time.Duration) {
	RedisClient.Set(ctx, token, userId, ttl)
}

func GetUserByToken(ctx context.Context, token string) (uint32, error) {
	result, err := RedisClient.Get(ctx, token).Result()
	if err != nil {
		return 0, err
	}
	userId, _ := strconv.ParseUint(result, 10, 32)

	return uint32(userId), nil
}

func SetJTI(ctx context.Context, userId uint32, jti string) {
	err := RedisClient.SAdd(ctx, "jti_list:"+strconv.Itoa(int(userId)), jti).Err()
	if err != nil {
		klog.Error(err)
	}
}

func ListJTIList(ctx context.Context, userId uint32) ([]string, error) {
	return RedisClient.SMembers(ctx, "jti_list:"+strconv.Itoa(int(userId))).Result()
}

func AddJTI2BlackListed(ctx context.Context, jti string) {
	err := RedisClient.SAdd(ctx, "blacklisted:jti", jti).Err()
	if err != nil {
		klog.Error(err)
	}
}

func AddJTIList2BlackListed(ctx context.Context, jtiList []string) {
	err := RedisClient.SAdd(ctx, "blacklisted:jti", jtiList).Err()
	if err != nil {
		klog.Error(err)
	}
}

func CheckJITIsBlackListed(ctx context.Context, jti string) (bool, error) {
	return RedisClient.SIsMember(ctx, "blacklisted:jti", jti).Result()
}
