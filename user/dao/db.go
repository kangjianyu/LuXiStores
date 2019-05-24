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
}

type dbimpl struct {
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
	if err := common.MysqlClient.Table(tableName).Where("uid=?", uid).First(&userInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userInfo, nil
		}
		return userInfo, err
	}
	return userInfo, nil
}
