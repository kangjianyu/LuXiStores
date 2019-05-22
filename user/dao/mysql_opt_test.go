package dao

import (
	"fmt"
	"testing"
)

func TestFetchUser(t *testing.T) {
	Init()
	userinfo := UserInfo{UserName: "kjy"}
	ret, err := FetchUser(&userinfo)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	fmt.Println(ret)
}
