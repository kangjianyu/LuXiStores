package k_client_test

import (
	"LuXiStores/common"
	goods_dao "LuXiStores/goods/dao"
	"fmt"
	"testing"
)

//func TestMysqlClient_GetByUid(t *testing.T) {
//	common.Init()
//	userprofile := &user_dao.UserProfile{}
//	DB := common.MysqlClient.GetByUid("userprofile",1,userprofile)
//	fmt.Println(DB,userprofile.Nick)
//
//}
//func TestMysqlClient_Ping(t *testing.T) {
//	common.Init()
//	common.MysqlClient.Ping()
//}
//func TestMysqlClient_UpdateByProfile(t *testing.T) {
//	common.Init()
//	err := common.MysqlClient.UpdateByProfile("userprofile",1,"sql修改",1997,0)
//	fmt.Println(err)
//}
//func TestMysqlClient_Insert(t *testing.T) {
//	common.Init()
//	info := user_dao.UserInfo{Username:"无值测试",
//
//	}
//	err := common.MysqlClient.InsertUserInfo("userinfo",info.Username,info.Password,info.Email,info.Status,info.Phone)
//	fmt.Println(err)
//}
//
//func TestMysqlClient_GetGoodsCategory(t *testing.T) {
//	common.Init()
//	info := goods_dao.GoodsCategory{}
//	_ = common.MysqlClient.GetGoodsCategory("category",1,&info)
//	fmt.Println(info.Name,"名字",info.Key)
//}
//func TestMysqlClient_GetGoodsCategoryNext(t *testing.T) {
//	common.Init()
//	infos := []goods_dao.GoodsCategory{}
//	_ = common.MysqlClient.GetGoodsCategoryNext("category","",1,2,&infos)
//	for i,x:= range infos{
//		fmt.Println(i,x)
//	}
//}
//
//func TestMysqlClient_UpdateByInfo(t *testing.T) {
//	common.Init()
//	ret := common.MysqlClient.UpdateUserInfo("userinfo",37,"222","1233","222")
//	fmt.Println(ret)
//}

func TestMysqlClient_GetGoodsCollectionByUid(t *testing.T) {
	common.Init()
	info := []goods_dao.GoodsInfoSlice{}
	ret :=common.MysqlClient.GetGoodsCollectionByUid("mmall_product","mmall_product_collection",1,&info)
	fmt.Println(ret.Error,ret.Value)
}

//func TestMysqlClient_GetGoodsInfo(t *testing.T) {
//	common.Init()
//	info := []goods_dao.GoodsInfo{}
//	count := 1
//	ret := common.MysqlClient.GetGoodsInfo("mmall_product",2,20,0,"stock",&info,&count)
//	for i,x := range info{
//		fmt.Println(i,x,ret)
//	}
//	fmt.Println(count)
//}