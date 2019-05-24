package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchUser(t *testing.T) {
	Init()
	userinfo := UserInfo{UserName: "kjy"}
	ret, err := FetchUserByName(&userinfo)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	fmt.Println(ret)
}
func TestInsertUser(t *testing.T) {
	Init()
	user := &UserInfo{UserName: "无手机", Password: "kang123", Email: "666.com", Gender: 2, CreateTime: time.Now()}
	err := InsertUser(user)
	if err != nil {
		fmt.Println("插入失败", err)
	}
	fmt.Println(err)
}
