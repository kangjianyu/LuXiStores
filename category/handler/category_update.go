package category_handler

import (
	category_dao "LuXiStores/category/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)
type CategoryUpdateData struct {
	CategoryId  uint64 	`json:"category_id"`
	Name 		string 	`json:"name"`
	SortOrder	uint64  `json:"sort_order"`
	ParentId    uint64 	`json:"parent_id"`
	Status 		uint8	`json:"status"`
	Level       uint64  `json:"level"`
	Key 		string  `json:"key"`
}

func CategoryUpdate(c *gin.Context){
	prefix := "CategoryUpdate"
	data,err:=ioutil.ReadAll(c.Request.Body)
	Data := CategoryUpdateData{}
	err = json.Unmarshal(data,&Data)
	if err!=nil||Data.Name==""{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = category_dao.DB.UpdateGoodsType(Data.CategoryId,Data.Name,Data.ParentId,Data.Status,Data.SortOrder,Data.Key,Data.Level)
	if err!=nil{
		log.Warn(prefix,"node update error:%v,node:%v",err,Data)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed node data:%v",Data)
	common.BuildResp(c,nil,nil)
	return
}
