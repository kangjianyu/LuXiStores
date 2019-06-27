package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GetGoodsInfo(c *gin.Context){
	strstart := c.DefaultQuery("start","0")
	strcategoryId := c.Query("category_id")
	strcount := c.DefaultQuery("count","20")
	sortOrder := c.DefaultQuery("sort","")
	categoryId,err := strconv.Atoi(strcategoryId)
	count,err := strconv.Atoi(strcount)
	start,err := strconv.Atoi(strstart)
	if err!=nil||start<0||categoryId<=0{
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
	common.BuildResp(c,data,nil)
	return
}


func GetGoodsInfoDetail(c *gin.Context){
	strid := c.Query("id")
	id,err :=strconv.Atoi(strid)
	if err!=nil||id<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	data,err := goods_dao.DB.GetGoodInfoDetail(uint64(id))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	common.BuildResp(c,data,nil)
	return


}
