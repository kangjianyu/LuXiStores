package dao

import (
	"time"
)
import log "github.com/jeanphorn/log4go"

type UserInfo struct {
	UId        uint64    `json:"uid"`
	UserName   string    `json:"username"`
	PassWord   string    `json:"password"`
	Email      string    `json:"email"`
	Gender     uint8     `json:"gender"`
	Status     uint8     `json:"status"`
	Phone      string    `json:"phone"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

//InsertUser 注册用户
func InsertUser(userinfo *UserInfo) (err error) {
	const prefix = "InsertUser"
	userinfo.CreateTime = time.Now()
	_, err = insert("INSERT INTO userinfo VALUES(?,?,?,?,?,?,?,NULL)",
		userinfo.UId, userinfo.UserName, userinfo.PassWord,
		userinfo.Email, userinfo.Gender, userinfo.Phone, userinfo.CreateTime)
	if err != nil {
		log.Error(prefix+"insertUser error: %v, : %v", err, userinfo)
		return
	}
	log.Info(prefix+"result userinfo: %+v", userinfo)
	return
}

//FetchUser 查询单个用户
func FetchUser(userinfo *UserInfo) (ret map[string]string, err error) {
	const prefix = "FetchUser"
	ret, err = FetchRow("SELECT uid, password, email, gender,status FROM userinfo where username=?", userinfo.UserName)
	if err != nil {
		log.Error(prefix+" error: %v, : %v", err, userinfo)
		return
	}
	log.Info(prefix+"result userinfo: %+v", ret)
	return
}

//UpdateUser
func UpDateUser(userinfo *UserInfo) (err error) {

	return
}
