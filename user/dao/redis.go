package user_dao

import (
	"LuXiStores/common"
	"time"
)

var Rds iRedis = redisImpl{}

type iRedis interface {
	GetUserToken(key string) (value string, err error)
	SetUserToken(key string, value string, TTl time.Duration) (err error)
	DelUserToken(key string) (err error)
}

type redisImpl struct{}

func (redisImpl) GetUserToken(key string) (value string, err error) {
	return common.RedisClient.Get(key).Result()
}

func (redisImpl) SetUserToken(key string, value string, TTl time.Duration) (err error) {
	_, err = common.RedisClient.Set(key, value, TTl).Result()
	return
}

func (redisImpl) DelUserToken(key string) (err error) {
	_, err = common.RedisClient.Del(key).Result()
	return
}
