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

//商品收藏
func (m *MysqlClient) InsertGoodsCollection(table string,uid int64,productId int64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`uid`,`product_id`) VALUES (%d,%d)",table,uid,productId)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func(m *MysqlClient) DelGoodsCollection(table string,uid int64,productId int64)*gorm.DB{
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `uid` = %d AND `product_id` = %d",table,uid,productId)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

func(m *MysqlClient) GetGoodsCollectionByUid(goodsTable string,collectionTable string,uid int64,offset int64,limit int64,user interface{}) *gorm.DB{
	//query := fmt.Sprintf("SELECT `id`,`name`,`main_image`,`price`,`stock`,`status` FROM `%s` WHERE id IN ",goodsTable)
	//subquery := fmt.Sprintf("(SELECT `product_id` FROM `%s` WHERE uid=%d)",collectionTable,uid)
	//sql := query+subquery
	ret := m.DB.LogMode(true).Table(goodsTable).Where("id in (?)",m.DB.Table(collectionTable).Select("product_id").Where("uid=?",uid).QueryExpr()).Limit(limit).Offset(offset).Find(user)
	return ret
}

//短链接
func (m *MysqlClient) InsertShortUrl(table string,shorturl string,longurl string)*gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` ( `short_url`,`long_url`) VALUES ('%s','%s')",table,shorturl,longurl)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) GetLongUrl(table string,shorturl string,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("short_url=?",shorturl).First(user)
	return ret
}

//func (m *MysqlClient) GetMaxShortId(table string,shorturl string)*gorm.DB{
//	ret := m.DB.Table(table).LogMode(true)
//
//}

//订单
func (m *MysqlClient) UpdateOrderStatus(table string,nowtime string) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `status`= %d WHERE `create_time` <= %s",table,4,nowtime)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) InsertOrderComment(table string,orderId,uid,start int64,context string)*gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` ( `order_id`,`uid`,`context`,`start`) VALUES (%d,%d,'%s',%d)",table,orderId,uid,context,start)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) DelOrderComment(table string,orderId,uid int64) *gorm.DB{
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `order_id` = %d AND `uid` = %d",table,orderId,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret

}
func(m *MysqlClient) GetOrderList(table string,uid uint64,count uint64,offset uint64,user interface{})*gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("user_id=?",uid).Limit(count).Offset(offset).Find(user)
	return ret
}

func (m *MysqlClient) GetOrderDetail(tablee string,orderId uint64,uid uint64,user interface{})*gorm.DB{
	ret := m.DB.Table(tablee).LogMode(true).Where("user_id=? AND id=?",uid,orderId).First(user)
	return ret
}

func (m *MysqlClient) InsertOrder(table string,tradeId string,orderId uint64,productId uint64,receiver string) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` ( `id`,`trade_id`,`product_id`,`address`) VALUES (%d,'%s',%d,'%s')",table,orderId,tradeId,productId,receiver)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

func (m *MysqlClient) GetMaxOrderId(table string,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Order("`id` desc ").Limit(1).First(user)
	return ret
}
//交易
func (m *MysqlClient) InsertTrade(table string,orderId int64, tradeId string, userId int64, productId int64, receiverId int64, price float64, amounts int64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`order_id`,`trade_id`,`uid`,`product_id`,`receiver_id`,`price`,`amount`) VALUES(%d,'%s',%d,%d,%d,%0.2f,%d)",table,orderId,tradeId,userId,productId,receiverId,price,amounts)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) UpdateOrderInfo(table string,orderId int64) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `status`=1 WHERE `id`=%d ",table,orderId)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
//购物车
func (m *MysqlClient) GetGoodsCartList(table string,uid uint64,count uint64,offset uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("user_id=?",uid).Limit(count).Offset(offset).Find(user)
	return ret
}
func (m *MysqlClient) UpdateGoodsCartCount(table string,productid uint64,quantity uint64,uid uint64)*gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `quantity`= %d WHERE `product_id`=%d AND `user_id`= %d",table,quantity,productid,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) AddGoodsCart(table string,uid uint64,quantity uint64,productId uint64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`user_id`,`product_id`,`quantity`) VALUES (%d,%d,%d) ON DUPLICATE  KEY UPDATE `quantity`=`quantity`+1",table,uid,productId,quantity)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

func (m *MysqlClient) DelGoodsCart(table string,uid uint64,productId uint64) *gorm.DB{
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `product_id`=%d AND `user_id`= %d ",table,productId,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}


//商品收货地址
func (m *MysqlClient) GetGoodsReceiverAddress( table string,uid uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("uid=?",uid).Find(user)
	return ret
}

func (m *MysqlClient) AddGoodsReceiverAddress(table string,uid uint64,nick string,tel string,mobile string,province string,city string,district string,address string,IsDefault uint8) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`uid`,`nick`,`tel`,`mobile`,`province`,`city`,`district`,`address`,`is_default`) VALUES (%d,'%s','%s','%s','%s','%s','%s','%s',%d) ",table,uid,nick,tel,mobile,province,city,district,address,IsDefault)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) DelGoodsReceiverAddress(table string,id uint64) *gorm.DB{
	sql := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = %d",table,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) UpdateGoodsReceiverAddress(table string,id uint64,nick string,tel string,mobile string,province string,city string,district string,address string) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `nick` =' %s',`tel` = '%s',`mobile`='%s',`province`='%s',`city`='%s',`district`='%s',`address`='%s' WHERE `id` = %d",table,nick,tel,mobile,province,city,district,address,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) ChangeDefaultGoodsReceiverAddress(table string,uid uint64) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `is_default` = 0 WHERE `uid`=%d ",table,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) SetDefaultGoodsReceiverAddress(table string,id uint64) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `is_default`= 1 WHERE `id`= %d ",table,id)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) GetDefaultGoodsReceiverAddress(table string,uid uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("`is_default` = 1 and `uid`=?",uid).First(user)
	return ret
}
//商品信息
func (m *MysqlClient) GetGoodsInfo(table string,categoryId int64,count int64,offset int64,sortorder string,users interface{},) *gorm.DB {
	ret := m.DB.Table(table).LogMode(true).Select("id,name,price,main_image,stock").Where("category_id=?",categoryId).Limit(count).Offset(offset).Order(sortorder).Find(users)
	return ret
}
func (m *MysqlClient) GetSomeGoodsInfoDetail(table string,user interface{},count int64,offset int64,productId ...int64)*gorm.DB{
	ret := m.DB.LogMode(true).Table(table).Where("id in (?)",productId).Limit(count).Offset(offset).Find(user)
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

func (m *MysqlClient) InsertUserInfo(table string,uid uint64,Username string,Password string,Email string,Status uint8,Phone string) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`uid`,`username`, `password`, `email`, `status`, `phone`) VALUES(%d,'%s','%s','%s',%d,'%s')",table,uid,Username,Password,Email,Status,Phone)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}

func (m *MysqlClient) GetMaxUid(table string,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Order("`uid` desc ").Limit(1).First(user)
	return ret
}
//用户资料
func (m *MysqlClient) InsertUserProfile(table string,uid uint64, nick string, birthDate uint64, gender uint8)*gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`uid`,`nick`,`birth_date`,`gender`) VALUES(%d,'%s',%d,%d)",table,uid,nick,birthDate,gender)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
func (m *MysqlClient) GetUserProfileByUid(table string,uid uint64,user interface{}) *gorm.DB {
	ret := m.DB.Table(table).Where("uid=?",uid).First(user)
	return ret
}
func (m *MysqlClient) UpdateUserProfile(table string,uid uint64,nick string,birthDate uint64,gender uint8) *gorm.DB{
	sql := fmt.Sprintf("UPDATE `%s` SET `nick`= '%s', `birth_date`= %d ,`gender`= %d WHERE uid = %d",table,nick,birthDate,gender,uid)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
//用户会员
func (m *MysqlClient) GetUserSuperInfo(table string,uid uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).LogMode(true).Where("uid=?",uid).First(user)
	return ret
}
func (m *MysqlClient) InsertUserSuperInfo(table string,uid uint64,StartTime int64,EndTime int64,renew int64) *gorm.DB{
	sql := fmt.Sprintf("INSERT INTO `%s` (`uid`,`start_time`,`end_time`) VALUES (%d,FROM_UNIXTIME( %d, '%s'),FROM_UNIXTIME( %d, '%s')) ON DUPLICATE  KEY UPDATE `end_time` = DATE_ADD(`end_time`, INTERVAL %d DAY) ",table,uid,StartTime,"%Y%m%d",EndTime,"%Y%m%d",renew)
	ret := m.DB.LogMode(true).Exec(sql)
	return ret
}
//func (m *MysqlClient) Ping() {
//	ret := m.DB.Table("userprofile")
//}
