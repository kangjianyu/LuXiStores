package user_dao

import (
	"LuXiStores/common"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB iDB = dbimpl{}

type iDB interface {
	AddUserInfo(info UserInfo) error
	GetUserInfoByUsername(username string) (UserInfo, error)
	GetUserInfoByUid(uid uint64) (UserInfo, error)
	GetUserProfileByUid(uid uint64) (UserProfile,error)
	UpdateUserProfile(uid uint64,nick string,birthDate uint64,gender uint8) error
	AddUserProfile(uid uint64,nick string,birthDate uint64,gender uint8) error
	UpdateUserInfo(uid uint64,password string,email string,phone string) error
	GetMaxUid() (uint64,error)
	GetUserSuperInfo(uid uint64) (UserSuper,error)
	AddUserSuperInfo(uid uint64,StartTime int64,EndTime int64,renew int64) error
}

type dbimpl struct {
}


//用户会员
func (dbimpl) GetUserSuperInfo(uid uint64) (UserSuper, error) {
	tablename := (&UserSuper{}).TableName()
	info := UserSuper{}
	ret := common.MysqlClient.GetUserSuperInfo(tablename,uid,&info)
	return info,ret.Error

}

func (dbimpl) AddUserSuperInfo(uid uint64, StartTime int64, EndTime int64, renew int64) error {
	tablename := (&UserSuper{}).TableName()
	ret := common.MysqlClient.InsertUserSuperInfo(tablename,uid,StartTime,EndTime,renew)
	return ret.Error

}

//用户资料

func (dbimpl) AddUserProfile(uid uint64, nick string, birthDate uint64, gender uint8) error {
	tablename := (&UserProfile{}).TableName()
	ret := common.MysqlClient.InsertUserProfile(tablename,uid,nick,birthDate,gender)
	return ret.Error

}
func (dbimpl) UpdateUserProfile(uid uint64,nick string,birthDate uint64,gender uint8) error {
	tablename := (&UserProfile{}).TableName()
	err := common.MysqlClient.UpdateUserProfile(tablename,uid ,nick,birthDate,gender).Error
	return err
}

func (dbimpl) GetUserProfileByUid(uid uint64) (UserProfile, error) {
	tableName := (&UserProfile{}).TableName()
	userprofile := UserProfile{}
	if err := common.MysqlClient.GetUserProfileByUid(tableName,uid,&userprofile).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err)==true{
			return userprofile,nil
		}
		return userprofile,err
	}
	return userprofile,nil
}
//用户信息
func (dbimpl) GetMaxUid() (uint64, error) {
	tablename := (&UserInfo{}).TableName()
	userinfo := UserInfo{}
	ret := common.MysqlClient.GetMaxUid(tablename,&userinfo)
	fmt.Println(userinfo.Uid,"最大id")
	return userinfo.Uid,ret.Error
}
func (dbimpl) UpdateUserInfo(uid uint64, password string,email string,phone string) error {
	tablename := (&UserInfo{}).TableName()
	err := common.MysqlClient.UpdateUserInfo(tablename,uid,password,email,phone).Error
	return err
}

func (dbimpl) AddUserInfo(info UserInfo) error {
	tableName := (&UserInfo{}).TableName()
	if err :=common.MysqlClient.InsertUserInfo(tableName,info.Uid,info.Username,info.Password,info.Email,info.Status,info.Phone).Error;err!=nil{
		return err
	}
	return nil
}

func (dbimpl) GetUserInfoByUsername(username string) (UserInfo, error) {
	tableName := (&UserInfo{}).TableName()
	userinfo := UserInfo{}
	if err := common.MysqlClient.GetUserInfoByUsername(tableName,username,&userinfo).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err)==true{
			return userinfo,nil
		}
		return userinfo,err
	}
	return userinfo,nil
}

func (dbimpl) GetUserInfoByUid(uid uint64) (UserInfo, error) {
	tableName := (&UserInfo{}).TableName()
	userInfo := UserInfo{}
	if err := common.MysqlClient.DB.Table(tableName).Where("uid=?", uid).First(&userInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userInfo, nil
		}
		return userInfo, err
	}
	return userInfo, nil
}
