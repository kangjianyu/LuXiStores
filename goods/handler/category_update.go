package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	data,err:=ioutil.ReadAll(c.Request.Body)
	Data := CategoryUpdateData{}
	err = json.Unmarshal(data,&Data)
	if err!=nil||Data.Name==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.UpdateGoodsType(Data.CategoryId,Data.Name,Data.ParentId,Data.Status,Data.SortOrder,Data.Key,Data.Level)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	common.BuildResp(c,nil,nil)
	return
}
