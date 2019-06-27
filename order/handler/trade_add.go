package order_handler

import (
	"LuXiStores/common"
	"LuXiStores/order/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"strings"
)



type AddTradeData struct {
	TradeId string `json:"trade_id"`
	Sign string `json:"sign"`
}

func AddTrade(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddTradeData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.TradeId==""||Data.Sign==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	check,err := order_dao.Rds.GetOrderId(Data.TradeId)
	if check!=Data.Sign{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	items := strings.Split(Data.TradeId,"_")
	user,product,price := "","",""
	amount,order,receiver := "","",""
	for _,x :=range items{
		if strings.Contains(x,"user"){
			index := strings.Index(x,"=")
			user = x[index+1:]
			fmt.Println(user,"id")
		}
		if strings.Contains(x,"product"){
			index := strings.Index(x,"=")
			product = x[index+1:]
			fmt.Println(product,"product")
		}
		if strings.Contains(x,"price"){
			index := strings.Index(x,"=")
			price = x[index+1:]
			fmt.Println(price,"price")
		}
		if strings.Contains(x,"amount"){
			index := strings.Index(x,"=")
			amount = x[index+1:]
			fmt.Println(amount,"amount")
		}
		if strings.Contains(x,"order"){
			index := strings.Index(x,"=")
			order = x[index+1:]
			fmt.Println(order,"order")
		}
		if strings.Contains(x,"receiver"){
			index := strings.Index(x,"=")
			receiver = x[index+1:]
			fmt.Println(receiver,"receiver")
		}

	}
	orderId,err := strconv.Atoi(order)
	userId,err := strconv.Atoi(user)
	productId,err := strconv.Atoi(product)
	receiverId,err := strconv.Atoi(receiver)
	prices,err := strconv.ParseFloat(price,10)
	amounts,err := strconv.ParseFloat(amount,10)
	amountPrice := prices*amounts
	err = order_dao.DB.AddTrade(int64(orderId),Data.TradeId,int64(userId),int64(productId),int64(receiverId),amountPrice,int64(amounts))
	err = order_dao.DB.OrderPayFinish(int64(orderId))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return

}
