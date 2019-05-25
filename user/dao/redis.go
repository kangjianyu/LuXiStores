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
	BlackListAdd(key string,field string)(err error)
	BlackListDel(key string,field ...string)(err error)
	BlackListCheck(key string,field...string)(value []interface{},err error)
}

type redisImpl struct{}

func (redisImpl) BlackListCheck(key string, field ...string) (value []interface{},err error) {
	return common.RedisClient.HMGet(key,field...).Result()
}

func (redisImpl) BlackListDel(key string, field ...string) (err error) {
	return common.RedisClient.HDel(key,field...).Err()
}

func (redisImpl) BlackListAdd(key string ,field string) (err error) {
	return common.RedisClient.Hash(key,field,1).Err()
}


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

