package receiver_handler

import (
	"LuXiStores/common"
	receiver_dao "LuXiStores/receiver/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)

type DelGoodsReceiverAddressData struct {
	Id 	uint64 `json:"id"`
	Uid uint64 `json:"uid"`
}
func DelGoodsReceiverAddress(c *gin.Context)  {
	prefix := "DelGoodsReceiverAddress"
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsReceiverAddressData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = receiver_dao.DB.DelGoodsReceiverAddress(Data.Id)
	if err!=nil{
		log.Warn(prefix,"")
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"del succeed id:%d,uid:%d",Data.Id,Data.Uid)
	common.BuildResp(c,nil,nil)
	return
}
