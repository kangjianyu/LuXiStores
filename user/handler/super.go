package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"io/ioutil"
	"strconv"
	"time"
)

func GetUserSuperInfo(c *gin.Context){
	strid := c.Query("uid")
	fmt.Println(strid)
	id,err := strconv.Atoi(strid)
	if id<=0||err!=nil{
		log.Warn("")
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	info,err := user_dao.DB.GetUserSuperInfo(uint64(id))
	if err!=nil{
		log.Warn("get super error %v uid:%d",err,id)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	log.Info("get superinfo succeed uid:%d",id)
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
		log.Warn("input data error %v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	nowtime := time.Now().Unix()
	endtime := time.Now().Unix()+(Data.PayDay*3600*24)
	err = user_dao.DB.AddUserSuperInfo(Data.Uid,nowtime,endtime,Data.PayDay)
	if err!=nil{
		log.Warn("insert super error %v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info("add superindo succeed uid:%d",Data.Uid)
	common.BuildResp(c,nil,nil)
	return

}
