package category_handler

import (
	"LuXiStores/common"
	"LuXiStores/category/dao"
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



