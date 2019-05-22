package dao

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"time"
)

func GetUserCookie(key string) (value string, err error) {
	client := redisClient
	result := client.Get(key)
	err = result.Err()
	value = result.Val()
	log.Info("get cookie in redis key:%s,value:%s", key, err)
	return
}

//SetValueByRedis 用户设置cookie
func SetUserCookie(key string, value string, TTl time.Duration) (err error) {
	client := redisClient
	fmt.Println(TTl * time.Second)
	err = client.Set(key, value, TTl*time.Second).Err()
	if err != nil {
		log.Error("set cookie error:%v in redis key:%s,value:%s", err, key, value)
		return
	}
	log.Info("set cookie in redis key:%s,value:%s", key, value)
	return
}
