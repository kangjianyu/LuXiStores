package receiver_handler

import (
	"LuXiStores/common"
	receiver_dao "LuXiStores/receiver/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)


type SetGoodsReceiverAddressData struct {
	Id uint64 			`json:"id"`
	Uid uint64 			`json:"uid"`
}
func SetGoodsReceiverAddress(c *gin.Context){
	prefix := "SetGoodsReceiverAddress"
	inData,err := ioutil.ReadAll(c.Request.Body)
	Data := SetGoodsReceiverAddressData{}
	err = json.Unmarshal(inData,&Data)
	if err!=nil||Data.Id<=0||Data.Uid<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = receiver_dao.DB.ChangeDefaultGoodsReceiverAddress(Data.Uid)
	err = receiver_dao.DB.SetDefaultGoodsReceiverAddress(Data.Id,Data.Uid)
	if err!=nil{
		log.Warn(prefix,"set default address error:%v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"set default address uid:%d,id:%d",Data.Uid,Data.Id)
	common.BuildResp(c,nil,nil)
	return
}
