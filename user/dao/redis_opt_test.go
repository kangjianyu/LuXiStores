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
	fmt.Println(time.Duration(times))
	err := SetUserCookie(Uuid, "kjy", time.Duration(times))
	fmt.Println(Uuid)
	if err != nil {
		fmt.Println("设置usercookie失败")
	}
	fmt.Println(Uuid)
}
func TestGetUserCookie(t *testing.T) {
	Init()
	value, err := GetUserCookie("2603b23f-f475-6fdb-83b9-f46c9dce1df9")
	fmt.Println(value, "结果")
	if err != nil {
		fmt.Println(value)
	}
	//fmt.Println(value)
}
