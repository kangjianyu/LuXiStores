package service

import (
	"LuXiStores/user/dao"
	log "github.com/jeanphorn/log4go"
	"time"
)

func SetCookieByRedis(uuid string, username string, ttl int64) (err error) {
	prefix := "SetCookieByRedis"
	err = dao.SetUserCookie(uuid, username, time.Duration(ttl))
	if err != nil {
		log.Error(prefix, "set Cookie error:%v", err)
		return
	}
	log.Info(prefix, "success set cookie:%s", uuid)
	return
}
func GetCookieByRedis(uuid string) (value string, err error) {
	prefix := "GetCookieByRedis"
	value, err = dao.GetUserCookie(uuid)
	if err != nil {
		log.Error(prefix, "get Cookie error:%v", err)
		return
	}
	log.Info(prefix, "success cookie:%s")
	return
}
