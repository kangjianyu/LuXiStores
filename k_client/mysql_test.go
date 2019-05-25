package k_client_test

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"fmt"
	"testing"
)

func TestMysqlClient_GetByUid(t *testing.T) {
	common.Init()
	userprofile := &user_dao.UserProfile{}
	DB := common.MysqlClient.GetByUid("userprofile",1,userprofile)
	fmt.Println(DB,userprofile.Nick)

}
func TestMysqlClient_Ping(t *testing.T) {
	common.Init()
	common.MysqlClient.Ping()
}
func TestMysqlClient_UpdateByProfile(t *testing.T) {
	common.Init()
	userprofile := &user_dao.UserProfile{Nick:"",Gender:0}
	err := common.MysqlClient.UpdateByProfile("userprofile",1,userprofile)
	fmt.Println(err)
}
