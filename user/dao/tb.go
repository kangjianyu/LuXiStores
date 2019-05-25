package user_dao

import "time"

type UserInfo struct {
	UId        uint64    `gorm:"column:uid" json:"uid"`
	UserName   string    `gorm:"column:username" json:"username"`
	PassWord   string    `gorm:"column:password" json:"password"`
	Email      string    `gorm:"column:email" json:"email"`
	Gender     uint8     `gorm:"column:gender" json:"gender"`
	Status     uint8     `gorm:"column:status" json:"status"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}
type UserProfile struct {
	Id 			uint64   `gorm:"column:id" json:"id"`
	UId			uint64 	 `gorm:"column:uid" json:"uid"`
	Username	string 	 `gorm:"column:username" json:"username"`
	Nick 		string	 `gorm:"column:nick" json:"nick"`
	Level 		uint64	 `gorm:"column:level" json:"level"`
	BirthDate 	uint64	 `gorm:"column:birth_date" json:"birth_date"`
	Gender		uint8  	 `gorm:"column:gender" json:"gender"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}
func (u *UserInfo) TableName() string {
	return "userinfo"
}
func(u *UserProfile) TableName() string{
	return "userprofile"
}

