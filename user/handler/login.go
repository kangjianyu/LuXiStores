package handler

import (
	"LuXiStores/user/service"
	log "github.com/jeanphorn/log4go"
)

//UserLogin 用户登录
func UserLogin(name string, password string) (valid bool, err error) {
	prefix := "UserLogin"
	userinfo, err := service.GetUserInfo(name)
	if err != nil {
		log.Error(prefix, "get user:%s error:%v ", name, err)
		return
	}
	if userinfo.PassWord == password {
		valid = true
		log.Info(prefix, "login successes user:%s", name)
		return
	} else {
		log.Info(prefix, "login error user:%s pwd:%s", name, password)
		return
	}
}

//SetCookie 设置用户cookie
func SetCookie(uuid string, username string) (err error) {
	prefix := "SetCookie"
	err = service.SetCookieByRedis(uuid, username, 10000)
	if err != nil {
		log.Error(prefix, "set cookie error:%v ", err)
		return
	}
	log.Info("set cookie in redis %s")
	return
}

//CheckCookie cookie验证
func CheckCookie(uuid string) (valid string, err error) {
	prefix := "CheckCookie"
	valid, err = service.GetCookieByRedis(uuid)
	if err != nil {
		log.Error(prefix, "check usercookie error:%v", err)
		return
	}
	log.Info(prefix, "check usercookie:%s success ", uuid)
	return
}
