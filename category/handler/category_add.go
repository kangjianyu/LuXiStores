package category_handler

import (
	category_dao "LuXiStores/category/dao"
	"LuXiStores/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)
type AddCategoryData struct {
	Name string  			`json:"name"`
	ParentId	int64		`json:"parent_id"`
}
func AddCategory(c *gin.Context){
	prefix := "AddCategory"
	data ,err := ioutil.ReadAll(c.Request.Body)
	categoryUpdateData := AddCategoryData{}
	err = json.Unmarshal(data,&categoryUpdateData)
	if err!=nil||categoryUpdateData.Name==""||categoryUpdateData.ParentId==0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	key,level,name,err := CategoryForNow(categoryUpdateData.ParentId)
	if err!=nil||name==""{
		log.Warn(prefix,"node is invalid error:%v,categoryId:%d",err)
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
		log.Warn(prefix,"add node error:%v,nodename:%d",err,categoryUpdateData.Name)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"")
	common.BuildResp(c,nil,nil)
	return

}
