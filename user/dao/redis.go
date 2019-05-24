package user_dao

import (
	"LuXiStores/common"
	log "github.com/jeanphorn/log4go"
	"time"
)

//GetUserCookie 获取用户token
func GetUserToken(key string) (value string, err error) {
	result := common.RedisClient.Get(key)
	err = result.Err()
	if err != nil {
		log.Error("get cookie error:%v in redis key:%s", err, key)
		return
	}
	value = result.Val()
	log.Info("get cookie in redis key:%s,value:%s", key, err)
	return
}

//SetValueByRedis 用户设置token
func SetUserToken(key string, value string, TTl time.Duration) (valid string, err error) {
	result := common.RedisClient.Set(key, value, TTl*time.Second)
	err = result.Err()
	valid = result.Val()
	if err != nil {
		log.Error("set cookie error:%v in redis key:%s,value:%s", err, key, value)
		return
	}
	log.Info("set cookie in redis key:%s,value:%s", key, value)
	return
}

//DelUserToken 删除用户token
func DelUserToken(key string) (valid int64, err error) {
	result := common.RedisClient.Del(key)
	valid = result.Val()
	err = result.Err()
	if err != nil {
		log.Error("del token error:%v in redis key:%s", err, key)
		return
	}
	log.Info("del token in redis key:%s", key)
	return
}
