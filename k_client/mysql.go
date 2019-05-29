package k_client

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/jinzhu/gorm"
	"runtime"
	"strings"
)

type MysqlClient struct {
	// 包装一层, 方便打日志和统计数据
	DB *gorm.DB

}
func NewMysqlClient(db *gorm.DB) *MysqlClient {
	return &MysqlClient{DB: db}
}

func MysqlLog(ret *gorm.DB) {
	pc, _, _, _ := runtime.Caller(2)
	detail := runtime.FuncForPC(pc)
	funcPath := strings.Split(detail.Name(), "/")
	callerName := funcPath[len(funcPath)-1]
	if ret.Error == nil {
		log.Info("(MYSQL %s)|%s|exist", callerName,ret.Value )
	} else if ret.RecordNotFound()==true {
		log.Info("(MYSQL %s)|%s|not exist", callerName,ret.Value )
	}else{
		log.Error("(REDIS %s)|%s|err:%s",callerName,ret.Error)
	}
}
//商品信息
func (m *MysqlClient) GetGoodsInfo(table string,categoryId int64,count int64,offset int64,sortorder string,users interface{},) *gorm.DB {
	ret := m.DB.Table(table).LogMode(true).Select("id,name,price,main_image,stock").Where("category_id=?",categoryId).Limit(count).Offset(offset).Order(sortorder).Find(users)
	return ret
}

func (m *MysqlClient) GetGoodsInfoDetail(table string,id uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("id=?",id).First(user)
	return ret
}
func (m *MysqlClient) InsertGoodsInfo(table string,categoryId uint64,name string,subtitle string,mainImage string,subImage string,detail string,price float64,stock uint64,status int64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`category_id`, `name`, `subtitle`, `main_image`, `sub_images`, `detail`, `price`, `stock`, `status`) VALUES ( %d,'%s','%s','%s','%s','%s',%.2f,%d,%d)",table,categoryId,name,subtitle,mainImage,subImage,detail,price,stock,status)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) UpdateGoodsInfo(table string,id uint64,categoryId uint64,name string,subtitle string,mainImage string,subImage string,detail string,price float64,stock uint64) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `category_id` = %d,`name` = '%s',`subtitle`= '%s',`main_image`='%s',`sub_images` = '%s', `detail`='%s',`price`= %.2f,`stock` = %d WHERE `id` = %d ",table,categoryId,name,subtitle,mainImage,subImage,detail,price,stock,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) UpdateGoodsStatus(table string,id uint64,status int64)*gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `status` = %d WHERE `id` = %d",table,status,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) DelGoodsInfo(table string,id uint64) *gorm.DB{
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = %d ",table,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

//商品类型
func (m *MysqlClient) GetGoodsCategoryNext(table string,key string,id int64,level int64,users interface{})*gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("`key` LIKE ? AND level=?",fmt.Sprintf("%s%d-%s",key,id,"%"),level).Find(users)
	return ret
}
func (m *MysqlClient) GetGoodsCategory(table string,id int64,user interface{})	*gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("id=?",id).First(user)
	return ret
}
func (m *MysqlClient) InsertGoodsCategory(table string,Name string,ParentId int64,Status int64,SortOrder int64,Key string,Level int64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`name`, `parent_id`, `status`, `sort_order`, `key`, `level`) VALUES ('%s',%d,%d,%d,'%s',%d )",table,Name,ParentId,Status,SortOrder,Key,Level)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) UpdateGoodsCategory(table string,id uint64, name string, parentId uint64, status uint8, sortOrder uint64, key string, level uint64) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `name` = '%s', `parent_id` = %d , `status`=%d , `sort_order` = %d , `key`='%s' , `level` = %d WHERE id = %d",table,name,parentId,status,sortOrder,key,level,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
//用户信息
func (m *MysqlClient) GetUserInfoByUsername(table string,username string,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("username=?",username).First(user)
	return ret
}

func (m *MysqlClient) UpdateUserInfo(table string,uid uint64,password string,email string,phone string) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `password` = '%s',`email` = '%s',`phone`= '%s' WHERE uid = %d",table,password,email,phone,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

func (m *MysqlClient) InsertUserInfo(table string,Username string,Password string,Email string,Status uint8,Phone string) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`username`, `password`, `email`, `status`, `phone`) VALUES('%s','%s','%s',%d,'%s')",table,Username,Password,Email,Status,Phone)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
//用户资料
func (m *MysqlClient) GetUserProfileByUid(table string,uid uint64,user interface{}) *gorm.DB {
	ret := m.DB.Table(table).Where("uid=?",uid).First(user)
	return ret
}
func (m *MysqlClient) UpdateUserProfile(table string,uid uint64,nick string,birthDate uint64,gender uint8) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `nick`= '%s', `birth_date`= %d ,`gender`= %d WHERE uid = %d",table,nick,birthDate,gender,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}


//func (m *MysqlClient) Ping() {
//	ret := m.DB.Table("userprofile")
//}
