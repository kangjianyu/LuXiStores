package user_dao

import "time"

type UserInfo struct {
	Uid        uint64    `gorm:"column:uid" json:"uid"`
	Username   string    `gorm:"column:username" json:"username"`
	Password   string    `gorm:"column:password" json:"password"`
	Email      string    `gorm:"column:email" json:"email"`
	Status     uint8     `gorm:"column:status" json:"status"`
	Phone      string    `gorm:"column:phone" json:"phone"`
}
type UserProfile struct {
	Id 			uint64   `gorm:"column:id" json:"id"`
	UId			uint64 	 `gorm:"column:uid" json:"uid"`
	Username	string 	 `gorm:"column:username" json:"username"`
	Nick 		string	 `gorm:"column:nick" json:"nick"`
	Level 		uint64	 `gorm:"column:level" json:"level"`
	BirthDate 	uint64	 `gorm:"column:birth_date" json:"birth_date"`
	Gender		uint8  	 `gorm:"column:gender" json:"gender"`
}

type UserSuper struct {
	Uid			uint64 	 `gorm:"column:uid" json:"uid"`
	StartTime 	time.Time   `gorm:"column:start_time" json:"start_time"`
	EndTime 	time.Time   `gorm:"column:end_time" json:"end_time"`
}

func (u *UserSuper) TableName() string{
	return "mmall_user_super"
}
func (u *UserInfo) TableName() string {
	return "userinfo"
}
func(u *UserProfile) TableName() string{
	return "userprofile"
}

