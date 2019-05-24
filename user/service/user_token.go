package service

import (
	"LuXiStores/user/dao"
	log "github.com/jeanphorn/log4go"
	"time"
)

func SetUserTokenByRedis(token string, username string, ttl int64) (valid string, err error) {
	prefix := "SetCookieByRedis"
	valid, err = dao.SetUserToken(token, username, time.Duration(ttl))
	if err != nil {
		log.Error(prefix, "set Cookie error:%v", err)
		return
	}
	log.Info(prefix, "success set cookie:%s", token)
	return
}
func GetUserTokenByRedis(token string) (value string, err error) {
	prefix := "GetCookieByRedis"
	value, err = dao.GetUserToken(token)
	if err != nil {
		log.Error(prefix, "get Cookie error:%v", err)
		return
	}
	log.Info(prefix, "success cookie:%s")
	return
}

func DelUserTokenByRedis(token string) (valid int64, err error) {
	prefix := "DelUserTokenByRedis"
	valid, err = dao.DelUserToken(token)
	if err != nil {
		log.Error(prefix, "del token error:%v", err)
		return
	}
	log.Info(prefix, "del token:%s", token)
	return
}
