package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
}

func TestFetchRow(t *testing.T) {
	db, err := gorm.Open("mysql", "root:KANG345876@qq.com@/luxistores?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	creattiem := time.Now()
	newuser := UserInfo{UserName: "新的连接", CreateTime: creattiem}
	valid := db.NewRecord(newuser)
	fmt.Println(valid)
	DB := db.Create(&newuser)
	DB.NewRecord(newuser)
}
