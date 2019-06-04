package category_dao

import "LuXiStores/common"

var DB iBD = dbimpl{ }
type iBD interface {
	GetGoodsTypeByidForNext(key string,categoryId int64,level int64,) ([]GoodsCategory,error)
	GetGoodsTypeByid(id int64)(GoodsCategory,error)
	UpdateGoodsType(id uint64,name string,parentId uint64,status uint8,sortOrder uint64,key string,level uint64) error
	AddGoodsType(info GoodsCategory) error


}

type dbimpl struct {

}

//商品类型
func (dbimpl) UpdateGoodsType(id uint64, name string, parentId uint64, status uint8, sortOrder uint64, key string, level uint64) error{
	tablename := (&GoodsCategory{}).TableName()
	err := common.MysqlClient.UpdateGoodsCategory(tablename,id,name,parentId,status,sortOrder,key,level).Error
	return err
}
func (dbimpl) AddGoodsType(info GoodsCategory) error {
	tablename := (&GoodsCategory{}).TableName()
	ret := common.MysqlClient.InsertGoodsCategory(tablename,info.Name,info.ParentId,info.Status,info.SortOrder,info.Key,info.Level)
	return ret.Error
}

func (dbimpl) GetGoodsTypeByid(id int64) (GoodsCategory, error) {
	tablename := (&GoodsCategory{}).TableName()
	goodCategory := GoodsCategory{}
	err := common.MysqlClient.GetGoodsCategory(tablename,id,&goodCategory).Error
	return goodCategory,err
}

func (dbimpl) GetGoodsTypeByidForNext(key string, categoryId int64, level int64,) ([]GoodsCategory,error) {
	tablename := (&GoodsCategory{}).TableName()
	goodsInfos := []GoodsCategory{}
	err := common.MysqlClient.GetGoodsCategoryNext(tablename,key,categoryId,level,&goodsInfos).Error
	return goodsInfos,err
}