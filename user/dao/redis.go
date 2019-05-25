package user_dao

import (
	"LuXiStores/common"
	log "github.com/jeanphorn/log4go"
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
	value, err = common.RedisClient.Get(key).Result()
	if err != nil {
		log.Error("get cookie error:%v in redis key:%s", err, key)
		return
	}
	log.Info("get cookie in redis key:%s,value:%s", key, err)
	return
}

func (redisImpl) SetUserToken(key string, value string, TTl time.Duration) (err error) {
	_, err = common.RedisClient.Set(key, value, TTl*time.Second).Result()
	if err != nil {
		log.Error("set cookie error:%v in redis key:%s,value:%s", err, key, value)
		return
	}
	log.Info("set cookie in redis key:%s,value:%s", key, value)
	return
}

func (redisImpl) DelUserToken(key string) (err error) {
	_, err = common.RedisClient.Del(key).Result()
	if err != nil {
		log.Error("del token error:%v in redis key:%s", err, key)
		return
	}
	log.Info("del token in redis key:%s", key)
	return
}
