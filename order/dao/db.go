package order_dao

import "LuXiStores/common"

var DB iBD = dbimpl{}

type iBD interface {
	AddOrder(tradeId string) error
	//AddTrade()
}


type dbimpl struct {

}

func (dbimpl) AddOrder(tradeId string) error {
	tablename := (&OrderInfo{}).TableName()
	ret := common.MysqlClient.InsertOrder(tablename,tradeId)
	return ret.Error
}
