package receiver_handler

import (
	"LuXiStores/common"
	receiver_dao "LuXiStores/receiver/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)

type UpdateGoodsReceiverAddressData struct {
	Id 		   uint64 		`json:"id"`
	Uid        uint64 		`json:"uid"`
	Nick       string 		`json:"nick"`
	Tel        string 		`json:"tel"`
	Mobile     string 		`json:"mobile"`
	Province   string 		`json:"province"`
	City       string 		`json:"city"`
	District   string 		`json:"district"`
	Address    string 		`json:"address"`
}

func UpdateGoodsReceiverAddress(c *gin.Context){
	prefix:="UpdateGoodsReceiverAddress"
	inData,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsReceiverAddressData{}
	err = json.Unmarshal(inData,&Data)
	if err!=nil||Data.Id<=0||Data.Uid<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = receiver_dao.DB.UpdateGoodsReceiverAddress(Data.Id,Data.Uid,Data.Nick,Data.Tel,Data.Mobile,Data.Province,Data.City,Data.District,Data.Address)
	if err!=nil{
		log.Warn(prefix,"update receiver address error %v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"update receiver address succeed id:%d,",Data.Id,)
	common.BuildResp(c,nil,nil)
	return
}
