package category_handler

import (
	category_dao "LuXiStores/category/dao"
	"LuXiStores/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)
type CategoryAddData struct {
	Name string  			`json:"name"`
	ParentId	int64		`json:"parent_id"`
}
func CategoryAdd(c *gin.Context){
	data ,err := ioutil.ReadAll(c.Request.Body)
	categoryUpdateData := CategoryAddData{}
	err = json.Unmarshal(data,&categoryUpdateData)
	if err!=nil||categoryUpdateData.Name==""||categoryUpdateData.ParentId==0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	key,level,name,err := CategoryForNow(categoryUpdateData.ParentId)
	if err!=nil||name==""{
		common.BuildResp(c,nil,errors.New("父节点无效"))
		return
	}
	info := category_dao.GoodsCategory{
		Name:       categoryUpdateData.Name,
		ParentId:   categoryUpdateData.ParentId,
		Status:     0,
		Key:        fmt.Sprintf("%s%d-",key,categoryUpdateData.ParentId),
		Level:      level+1,
	}
	err = category_dao.DB.AddGoodsType(info)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return

}
