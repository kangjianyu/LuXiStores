package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func GetUserSuperInfo(c *gin.Context){
	strid := c.Query("uid")
	fmt.Println(strid)
	id,err := strconv.Atoi(strid)
	if id<=0||err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	info,err := user_dao.DB.GetUserSuperInfo(uint64(id))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	common.BuildResp(c,info,nil)
	return
}

type AddUserSuperInfoData struct {
	Uid uint64 `json:"uid"`
	PayDay int64 `json:"pay_day"`
	TradeId string `json:"trade_id"`
}
func AddUserSuperInfo(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddUserSuperInfoData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Uid<=0||Data.PayDay<=0||Data.TradeId==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	nowtime := time.Now().Unix()
	endtime := time.Now().Unix()+(Data.PayDay*3600*24)
	err = user_dao.DB.AddUserSuperInfo(Data.Uid,nowtime,endtime,Data.PayDay)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return

}
