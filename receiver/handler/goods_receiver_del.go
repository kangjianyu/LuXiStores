package receiver_handler

import (
	"LuXiStores/common"
	receiver_dao "LuXiStores/receiver/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type DelGoodsReceiverAddressData struct {
	Id 	uint64 `json:"id"`
	Uid uint64 `json:"uid"`
}
func DelGoodsReceiverAddress(c *gin.Context)  {
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsReceiverAddressData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = receiver_dao.DB.DelGoodsReceiverAddress(Data.Id)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}
