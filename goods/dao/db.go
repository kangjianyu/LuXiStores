package goods_dao

import (
	"LuXiStores/common"
)

var DB iBD = dbimpl{}

type iBD interface {
	GetGoodsInfo(categoryId int64,count int64,page int64,sortorder string) ([]GoodsInfoSlice,error)
	GetGoodInfoDetail(goodsId uint64)(GoodsInfo,error)
	AddGoodsInfo(categoryId uint64,name string,subtitle string,mainImage string,subImages string,detail string,price float64,Stock uint64,status int64) error
	UpdateGoodsInfo(id uint64,categoryId uint64,name string,subtitle string,mainImage string,subImages string,detail string,price float64,Stock uint64) error
	UpdateGoodsStatus(id uint64,status int64) error
	DelGoodsInfo(id uint64)error
	AddGoodsCollection(uid int64,productId int64) error
	DelGoodsCollection(uid int64,productId int64) error
	GetGoodsCollectionByUid(uid int64) ([]GoodsCollection,error)
	GetSomeGoodsCollection(count int64,offset int64,productId ...int64) ([]GoodsInfoSlice,error)
}


type dbimpl struct {

}

func (dbimpl) GetSomeGoodsCollection(count int64,offset int64,productId ...int64) ([]GoodsInfoSlice, error) {
	tablename := (&GoodsInfo{}).TableName()
	infos := []GoodsInfoSlice{}
	ret := common.MysqlClient.GetSomeGoodsInfoDetail(tablename,&infos,count,offset,productId...)
	return infos,ret.Error
}

func (dbimpl) AddGoodsCollection(uid int64, productId int64) error {
	tablename := (&GoodsCollection{}).TableName()
	ret := common.MysqlClient.InsertGoodsCollection(tablename,uid,productId)
	return ret.Error
}

func (dbimpl) DelGoodsCollection(uid int64, productId int64) error {
	tablename := (&GoodsCollection{}).TableName()
	ret := common.MysqlClient.DelGoodsCollection(tablename,uid,productId)
	return ret.Error

}

func (dbimpl) GetGoodsCollectionByUid(uid int64) ([]GoodsCollection, error) {
	tablename := (&GoodsCollection{}).TableName()
	info := []GoodsCollection{}
	ret := common.MysqlClient.GetGoodsCollectionByUid(tablename,uid,&info)
	return info,ret.Error
}

//商品信息
func (dbimpl) GetGoodsInfo(categoryId int64, count int64, page int64, sortorder string) ([]GoodsInfoSlice,error) {
	tablename := (&GoodsInfo{}).TableName()
	infos := []GoodsInfoSlice{}
	err := common.MysqlClient.GetGoodsInfo(tablename,categoryId,count,page,sortorder,&infos,).Error
	return infos,err
}
func (dbimpl) GetGoodInfoDetail(goodsId uint64) (GoodsInfo, error) {
	tablename := (&GoodsInfo{}).TableName()
	infos := GoodsInfo{}
	err := common.MysqlClient.GetGoodsInfoDetail(tablename,goodsId,&infos).Error
	return infos,err
}

func (dbimpl) AddGoodsInfo(categoryId uint64, name string, subtitle string, mainImage string, subImages string, detail string, price float64, stock uint64,status int64) error{
	tablename := (&GoodsInfo{}).TableName()
	err :=common.MysqlClient.InsertGoodsInfo(tablename,categoryId , name , subtitle , mainImage , subImages , detail , price , stock,status ).Error
	return err
}

func (dbimpl) UpdateGoodsInfo(id uint64, categoryId uint64, name string, subtitle string, mainImage string, subImages string, detail string, price float64, Stock uint64) error {
	tablename := (&GoodsInfo{}).TableName()
	err := common.MysqlClient.UpdateGoodsInfo(tablename,id,categoryId,name,subtitle,mainImage,subImages,detail,price,Stock).Error
	return err
}

func (dbimpl) UpdateGoodsStatus(id uint64, status int64) error {
	tablename := (&GoodsInfo{}).TableName()
	err := common.MysqlClient.UpdateGoodsStatus(tablename,id,status).Error
	return err

}
func (dbimpl) DelGoodsInfo(id uint64) error {
	tablename := (&GoodsInfo{}).TableName()
	err := common.MysqlClient.DelGoodsInfo(tablename,id).Error
	return err
}








