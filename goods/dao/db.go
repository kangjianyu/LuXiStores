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

}


type dbimpl struct {

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








