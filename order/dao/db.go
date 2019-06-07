package order_dao

import "LuXiStores/common"

var DB iBD = dbimpl{}

type iBD interface {
	AddOrder(tradeId string,orderId uint64,productId uint64) error
	GetMaxOrderId() (uint64,error)
	AddTrade(orderId int64,tradeId string,userId int64,productId int64,receiverIdd int64,price float64,)
	AddOrderComment(orderId int64,Uid int64,start int64,context string) error
	DelOrderComment(orderId int64,uid int64) error
}


type dbimpl struct {

}

func (dbimpl) DelOrderComment(orderId int64, uid int64) error {
	tablename := (&OrderComment{}).TableName()
	ret := common.MysqlClient.DelOrderComment(tablename,orderId,uid)
	return ret.Error
}

func (dbimpl) AddOrderComment(orderId int64, Uid int64, start int64, context string) error{
	tablename := (&OrderComment{}).TableName()
	ret := common.MysqlClient.InsertOrderComment(tablename,orderId,Uid,start,context).Error
	return ret
}

func (dbimpl) AddTrade(orderId int64, tradeId string, userId int64, productId int64, receiverIdd int64, price float64, ) {
	panic("implement me")
}

func (dbimpl) GetMaxOrderId() (uint64, error) {
	tablename := (&OrderInfo{}).TableName()
	user := OrderInfo{}
	ret := common.MysqlClient.GetMaxOrderId(tablename,&user)
	return user.Id,ret.Error
}

func (dbimpl) AddOrder(tradeId string,orderId uint64,productId uint64) error {
	tablename := (&OrderInfo{}).TableName()
	ret := common.MysqlClient.InsertOrder(tablename,tradeId,orderId,productId)
	return ret.Error
}
