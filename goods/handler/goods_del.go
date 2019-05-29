package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type DelGoodsInfoData struct {
	Id uint64	`json:"id"`
}

func DelGoodsInfo(c *gin.Context){
	InData,err :=ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsInfoData{}
	err = json.Unmarshal(InData,&Data)
	if err!=nil||Data.Id<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.DelGoodsInfo(Data.Id)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}
