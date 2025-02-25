package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/2/23 21:27
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: 薄雾算法
 */
/**
 * 1      2                                                     48         56       64
 * +------+-----------------------------------------------------+----------+----------+
 * |retain| increase                                            | salt2    |  salt1   |
 * +------+-----------------------------------------------------+----------+----------+
 * |0     | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
 * +------+-----------------------------------------------------+------------+--------+
 */

const (
	salt1Bit      = 8
	salt2Bit      = 8
	increaseBit   = 47
	salt2Shift    = salt1Bit
	increaseShift = salt2Shift + salt2Bit
	salt1Max      = (1 << salt1Bit) - 1
	salt2Max      = (1 << salt2Bit) - 1
	increaseMax   = (1 << increaseBit) - 1
)

type Mist struct {
	client     *redis.Client
	idRedisKey string
}

func NewMist(client *redis.Client, idRedisKey string) *Mist {
	return &Mist{client: client, idRedisKey: idRedisKey}
}

func (mist Mist) GenerateId() (id uint64, err error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	increaseNum, err := mist.client.Incr(context.Background(), mist.idRedisKey).Result()
	if err != nil {
		return 0, err
	}
	if increaseNum > increaseMax {
		panic("id generate : increaseNum > increaseMax")
	}

	salt1Num := r.Intn(salt1Max)
	salt2Num := r.Intn(salt2Max)
	id = (uint64(increaseNum) << increaseShift) | (uint64(salt2Num) << salt2Shift) | (uint64(salt1Num))

	return id, nil

}
