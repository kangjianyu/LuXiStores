package user_dao

import (
	"LuXiStores/common"
	"github.com/jinzhu/gorm"
)

var DB iDB = dbimpl{}

type iDB interface {
	AddUser(info UserInfo) error
	GetUserInfoByUsername(username string) (UserInfo, error)
	GetUserInfoByUid(uid uint64) (UserInfo, error)
	GetUserProfileByUid(uid uint64) (UserProfile,error)
	UpdateUserProfile(uid uint64,nick string,gender uint8) error
}

type dbimpl struct {
}

func (dbimpl) UpdateUserProfile(uid uint64, nick string,gender uint8) error {
	tablename := (&UserProfile{}).TableName()
	userrrofile := UserProfile{
		Nick:nick,
		Gender:gender,
	}
	err := common.MysqlClient.UpdateByProfile(tablename,uid,&userrrofile).Error
	return err
}

func (dbimpl) GetUserProfileByUid(uid uint64) (UserProfile, error) {
	tablename := (&UserProfile{}).TableName()
	userprofile := UserProfile{}
	if err := common.MysqlClient.GetByUid(tablename,uid,&userprofile).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err)==true{
			return userprofile,nil
		}
		return userprofile,err
	}
	return userprofile,nil
}

func (dbimpl) AddUser(info UserInfo) error {
	panic("implement me")
}

func (dbimpl) GetUserInfoByUsername(username string) (UserInfo, error) {
	panic("implement me")
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
