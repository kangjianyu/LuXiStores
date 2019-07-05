package backed

import (
	"LuXiStores/common"
	"context"
	"fmt"
	"time"
)

func InitBackend(){
	ctx := context.TODO()
	go RenewTicker(ctx)
	return
}

func RenewTicker2(ctx context.Context){
	ticker:= time.NewTicker(1*time.Minute)
	NotifyOrderTimeout(ctx,1,2)
	for range ticker.C{
		NotifyOrderTimeout(ctx,1,2)
		break
	}
}
func NotifyOrderTimeout(ctx context.Context,productId int,stock int){
	fmt.Println("我做到了")
}

func RenewTicker(ctx context.Context){
	ticker := time.NewTicker(1*time.Hour)
	NotifyRelationRenewProcess(ctx)
	for range ticker.C{
		NotifyRelationRenewProcess(ctx)
	}
}
func NotifyRelationRenewProcess(ctx context.Context){
	timestamp := time.Now().Unix()-(15*60)
	fmt.Println(time.Unix(timestamp,0).Format("20060102150405"))
	err := common.MysqlClient.GoodsTrackStock("mmall_order","mmall_product",	time.Unix(timestamp,0).Format("20060102150405"))
	if err!=nil{

	}
}

