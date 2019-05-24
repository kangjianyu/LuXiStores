package service

import (
	"LuXiStores/user/dao"
	log "github.com/jeanphorn/log4go"
	"strconv"
)

func GetUserInfoByName(name string) (userinfo dao.UserInfo, err error) {
	prefix := "GetUserInfo"
	user := &dao.UserInfo{UserName: name}
	ret, err := dao.FetchUserByName(user)
	userinfo.UserName = ret["username"]
	uid, err := strconv.Atoi(ret["uid"])
	if err != nil {
		log.Error(prefix, "uid:%s change error:%v ", ret["uid"], err)
		return
	}
	userinfo.UId = uint64(uid)
	userinfo.PassWord = ret["password"]
	gender, err := strconv.Atoi(ret["gender"])
	if err != nil {
		log.Error(prefix, "gender:%s change error:%v ", ret["gender"], err)
		return
	}
	userinfo.Gender = uint8(gender)
	status, err := strconv.Atoi(ret["status"])
	if err != nil {
		log.Error(prefix, "status:%s change error:%v ", ret["gender"], err)
		return
	}
	userinfo.Status = uint8(status)
	userinfo.Email = ret["email"]
	log.Info(prefix, "get userinfo%v", userinfo)
	return
}
func AddUserInfo(username string, password string, email string, phone string, gender string, arg ...string) (code int64, err error) {
	prefix := "AddUserInfo"
	intgender, err := strconv.Atoi(gender)
	code = 0
	if err != nil {
		code = 1001
		log.Error(prefix, "change gender error:%v", err)
		return
	}
	userinfo := &dao.UserInfo{
		UserName: username,
		PassWord: password,
		Email:    email,
		Phone:    phone,
		Gender:   uint8(intgender),
	}
	err = dao.InsertUser(userinfo)
	if err != nil {
		code = 1002
		log.Error(prefix, "get userinfo error:%v", err)
		return
	}
	return
}
func UpdateUserInfo() {

}
