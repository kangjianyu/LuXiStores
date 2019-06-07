package backed

import (
	"LuXiStores/common"
	"context"
	"time"
)

func InitBackend(){
	ctx := context.TODO()
	go RenewTicker(ctx)
	return
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
	common.MysqlClient.UpdateOrderStatus("mmall_order",	time.Unix(timestamp,0).Format("20060102150405"))

}