package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertUser(t *testing.T) {
	Init()
	user := &UserInfo{UId: 3, UserName: "kangjianyu", PassWord: "kang345876", Email: "666.com", Gender: 2, CreateTime: time.Now()}
	err := InsertUser(user)
	if err != nil {
		fmt.Println("插入失败", err)
	}
	fmt.Println(err)
}
