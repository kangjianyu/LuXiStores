package receiver_dao

import "LuXiStores/common"

var DB iBD = dbimpl{}

type iBD interface {
	GetGoodsReceiverAddress(uid uint64)([]GoodsReceiverAddress,error)
	AddGoodsReceiverAddress(uid uint64,nick string,tel string,mobile string,province string,city string,district string,address string,IsDefault uint8 ) error
	DelGoodsReceiverAddress(id uint64)error
	UpdateGoodsReceiverAddress(id uint64,uid uint64,nick string,tel string,mobile string,province string,city string,district string,address string) error
	ChangeDefaultGoodsReceiverAddress(uid uint64) error
	SetDefaultGoodsReceiverAddress(id uint64,uid uint64) error
	GetDefaultGoodsReceiverAddress(uid uint64)(GoodsReceiverAddress,error)

}


type dbimpl struct {

}

//商品收货地址
func (dbimpl) GetDefaultGoodsReceiverAddress(uid uint64) (GoodsReceiverAddress, error) {
	tablename := (&GoodsReceiverAddress{}).TableName()
	info := GoodsReceiverAddress{}
	err := common.MysqlClient.GetDefaultGoodsReceiverAddress(tablename,uid,&info).Error
	return info,err
}

func (dbimpl) AddGoodsReceiverAddress(uid uint64, nick string, tel string, mobile string, province string, city string, district string, address string, IsDefault uint8) error {
	tablename := (&GoodsReceiverAddress{}).TableName()
	err := common.MysqlClient.AddGoodsReceiverAddress(tablename,uid,nick,tel,mobile,province,city,district,address,IsDefault).Error
	return err
}

func (dbimpl) DelGoodsReceiverAddress(id uint64) error {
	tablename := (&GoodsReceiverAddress{}).TableName()
	err := common.MysqlClient.DelGoodsReceiverAddress(tablename,id).Error
	return err
}

func (dbimpl) UpdateGoodsReceiverAddress(id uint64, uid uint64, nick string, tel string, mobile string, province string, city string, district string, address string) error {
	tablename := (&GoodsReceiverAddress{}).TableName()
	err := common.MysqlClient.UpdateGoodsReceiverAddress(tablename,id,nick,tel,mobile,province,city,district,address).Error
	return err
}
func (dbimpl) ChangeDefaultGoodsReceiverAddress(uid uint64) error {
	tablename := (&GoodsReceiverAddress{}).TableName()
	err := common.MysqlClient.ChangeDefaultGoodsReceiverAddress(tablename,uid).Error
	return err
}
func (dbimpl) SetDefaultGoodsReceiverAddress(id uint64, uid uint64) error {
	tablename := (&GoodsReceiverAddress{}).TableName()
	err := common.MysqlClient.SetDefaultGoodsReceiverAddress(tablename,id).Error
	return err
}


func (dbimpl) GetGoodsReceiverAddress(uid uint64) ([]GoodsReceiverAddress, error) {
	tablename := (&GoodsReceiverAddress{}).TableName()
	info := []GoodsReceiverAddress{}
	err := common.MysqlClient.GetGoodsReceiverAddress(tablename,uid,&info).Error
	return info,err
}

