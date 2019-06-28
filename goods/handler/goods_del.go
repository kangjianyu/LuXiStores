package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)

type DelGoodsInfoData struct {
	Id uint64	`json:"id"`
}

func DelGoodsInfo(c *gin.Context){
	prefix:="DelGoodsInfo"
	InData,err :=ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsInfoData{}
	err = json.Unmarshal(InData,&Data)
	if err!=nil||Data.Id<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.DelGoodsInfo(Data.Id)
	if err!=nil{
		log.Warn(prefix,"del error:%v,id:%d",err,Data.Id)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed id:%d",Data.Id)
	common.BuildResp(c,nil,nil)
	return
}
