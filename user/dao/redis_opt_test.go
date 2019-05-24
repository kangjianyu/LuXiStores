package dao

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"testing"
	"time"
)

func TestSetUserCookie(t *testing.T) {
	Init()
	Uuid, _ := uuid.GenerateUUID()
	times := int64(1000)
	valid, err := SetUserToken(Uuid, "kjy", time.Duration(times))
	if err != nil {
		fmt.Println("设置usercookie失败")
	}
	fmt.Println(Uuid)
	fmt.Println(valid, "值")
}
func TestGetUserCookie(t *testing.T) {
	Init()
	value, err := GetUserToken("2603b23f-f475-6fdb-83b9-f46c9dce1df9")
	fmt.Println(value, "结果")
	if err != nil {
		fmt.Println(value)
	}
	//fmt.Println(value)
}
func TestDelUserToken(t *testing.T) {
	Init()
	valid, _ := DelUserToken("2603b23f-f475-6fdb-83b9-f46c9dce1df9")
	if valid != 1 {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}
}
