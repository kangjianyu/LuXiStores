package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"strconv"
	"strings"
)

func GetGoodsInfo(c *gin.Context){
	prefix := "GetGoodsInfo"
	strstart := c.DefaultQuery("start","0")
	strcategoryId := c.Query("category_id")
	strcount := c.DefaultQuery("count","20")
	sortOrder := c.DefaultQuery("sort","")
	categoryId,err := strconv.Atoi(strcategoryId)
	count,err := strconv.Atoi(strcount)
	start,err := strconv.Atoi(strstart)
	if err!=nil||start<0||categoryId<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	sortlist := make([]string,2)
	if sortlist = strings.Split(sortOrder,"_");len(sortlist)>=2&&sortlist[1] == "desc"{
		sortOrder = sortlist[0]+" desc"
	}else{
		sortOrder = sortlist[0]
	}

	page := start
	infos,err := goods_dao.DB.GetGoodsInfo(int64(categoryId),int64(count),int64(page),sortOrder)
	if err!=nil{
		log.Warn(prefix,"get error:%v,categoryId:%d",err,categoryId)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	has_next :=false

	if (count)<=len(infos){
		has_next=true
	}
	offset := start
	if has_next==true{
		offset += count
	}
	data := gin.H{
		"goods":infos,
		"total_count":len(infos),
		"offset":offset,
		"has_next":has_next,
	}
	log.Info(prefix,"succeed categoryId:%d",categoryId)
	common.BuildResp(c,data,nil)
	return
}


func GetGoodsInfoDetail(c *gin.Context){
	prefix := "GetGoodsInfoDetail"
	strid := c.Query("id")
	id,err :=strconv.Atoi(strid)
	if err!=nil||id<=0{
		log.Warn(prefix,"input data error:%d",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	data,err := goods_dao.DB.GetGoodInfoDetail(uint64(id))
	if err!=nil{
		log.Warn(prefix,"get error:%v,id:%d",err,id)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed id:%d",id)
	common.BuildResp(c,data,nil)
	return


}
