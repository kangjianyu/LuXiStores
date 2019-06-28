package category_handler

import (
	"LuXiStores/category/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"strconv"
)



func CategoryForNext(c *gin.Context){
	prefix := "CategoryForNext"
	strcategoryId := c.Query("category_id")
	categoryId,err:=strconv.Atoi(strcategoryId)
	if categoryId==0||err!=nil{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	info,err := category_dao.Rds.GetCategoryInfo("category_"+strcategoryId)
	if info!="" && err==nil{
		infos := []category_dao.GoodsCategory{}
		err = json.Unmarshal([]byte(info),&infos)
		log.Info(prefix,"succeed categoryId:%d",categoryId)
		common.BuildResp(c,infos,nil)
		return
	}

	key,level,_,err:=CategoryForNow(int64(categoryId))
	if err!=nil{
		log.Warn(prefix,"node is invalid categoryId:%d",categoryId)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	categoryInfos,err :=category_dao.DB.GetGoodsTypeByidForNext(key,int64(categoryId),level+1)
	if err!=nil{
		log.Warn(prefix,"next node is valid categoryId:%d",categoryId)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	jsonBytes, err := json.Marshal(categoryInfos)
	err = category_dao.Rds.SetCategoryInfo("category_"+strcategoryId,string(jsonBytes))
	if err!=nil{
		log.Warn(prefix,"set redis error:%v,categoryId:%d",err,categoryId)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed categoryId:%d",categoryId)
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



