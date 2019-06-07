package order_handler

import (
	"LuXiStores/common"
	"LuXiStores/order/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

type AddOrderCommentData struct {
	OrderId int64 `json:"order_id"`
	Start 	int64 `json:"start"`
	Uid 	int64 `json:"uid"`
	Context string `json:"context"`
}
type DelOrderCommentData struct {
	Uid 	int64 `json:"uid"`
	OrderId int64 `json:"order_id"`

}
func AddOrderComment(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddOrderCommentData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.OrderId==0||Data.Uid==0||Data.Start<0||Data.Start>5{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	context := filtercontext(Data.Context)

	err = order_dao.DB.AddOrderComment(Data.OrderId,Data.Uid,Data.Start,context)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return

}

func DelOrderComment(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelOrderCommentData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.OrderId==0||Data.Uid==0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	err = order_dao.DB.DelOrderComment(Data.OrderId,Data.Uid)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}


func filtercontext(context string)string{
	char := []string{
		"草","你妈","死","屎","共产党","砍",
	}
	for _,x := range char{
		context = strings.ReplaceAll(context,x,"*")
	}
	return context




}