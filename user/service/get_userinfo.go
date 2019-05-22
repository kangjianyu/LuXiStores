package service

import (
	"LuXiStores/user/dao"
	log "github.com/jeanphorn/log4go"
	"strconv"
)

func GetUserInfo(name string) (userinfo dao.UserInfo, err error) {
	user := &dao.UserInfo{UserName: name}
	ret, err := dao.FetchUser(user)
	userinfo.UserName = ret["username"]

	uid, err := strconv.Atoi(ret["uid"])
	if err != nil {
		log.Error("uid:%s change error:%v ", ret["uid"], err)
		return
	}
	userinfo.UId = uint64(uid)
	userinfo.PassWord = ret["password"]
	gender, err := strconv.Atoi(ret["gender"])
	if err != nil {
		log.Error("gender:%s change error:%v ", ret["gender"], err)
		return
	}
	userinfo.Gender = uint8(gender)
	status, err := strconv.Atoi(ret["status"])
	if err != nil {
		log.Error("status:%s change error:%v ", ret["gender"], err)
		return
	}
	userinfo.Status = uint8(status)
	userinfo.Email = ret["email"]
	return
}
