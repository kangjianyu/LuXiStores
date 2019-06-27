package category_handler

import (
	"LuXiStores/category/dao"
	"LuXiStores/common"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)



func CategoryForNext(c *gin.Context){
	strcategoryId := c.Query("category_id")
	categoryId,err:=strconv.Atoi(strcategoryId)
	if categoryId==0||err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	info,err := category_dao.Rds.GetCategoryInfo("category_"+strcategoryId)
	if info!="" && err==nil{
		infos := []category_dao.GoodsCategory{}
		err = json.Unmarshal([]byte(info),&infos)
		common.BuildResp(c,infos,common.ErrRedisKeyNotExist)
		return
	}

	key,level,_,err:=CategoryForNow(int64(categoryId))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	categoryInfos,err :=category_dao.DB.GetGoodsTypeByidForNext(key,int64(categoryId),level+1)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	jsonBytes, err := json.Marshal(categoryInfos)
	err = category_dao.Rds.SetCategoryInfo("category_"+strcategoryId,string(jsonBytes))
	if err!=nil{
		fmt.Println(err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,categoryInfos,nil)
	return
}


func CategoryForNow(categoryId int64)(key string,level int64,name string,err error){
	categoryInfo,err := category_dao.DB.GetGoodsTypeByid(categoryId)
	if err!=nil{
		return
	}
	return categoryInfo.Key,categoryInfo.Level,categoryInfo.Name,err
}



