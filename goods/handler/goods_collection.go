package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

type AddGoodsCollectionData struct {
	Uid int64 `json:"uid"`
	ProductId int64 `json:"product_id"`
}
type DelGoodsCollectionData struct {
	Uid int64 `json:"uid"`
	ProductId int64 `json:"product_id"`
}


func AddGoodsCollection(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddGoodsCollectionData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.ProductId<=0||Data.Uid<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.AddGoodsCollection(Data.Uid,Data.ProductId)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	common.BuildResp(c,nil,nil)
	return
}

func DelGoodsCollection(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsCollectionData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Uid<=0||Data.ProductId<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.DelGoodsCollection(Data.Uid,Data.ProductId)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}

func GetGoodsCollectionList(c *gin.Context){
	uidstr := c.Query("uid")
	offsetstr := c.Query("offset")
	limitstr := c.Query("limit")
	uid,err := strconv.Atoi(uidstr)
	offset,err := strconv.Atoi(offsetstr)
	limit,err := strconv.Atoi(limitstr)


	if err!=nil||uid<=0||limit<0||offset<0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	if limit==0{
		limit = 20
	}

	info,err := goods_dao.DB.GetGoodsCollectionByUid(int64(uid))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	productids := []int64{}
	for _,x := range info{
		productids = append(productids,x.ProductId)
	}
	if len(productids)<1{
		common.BuildResp(c,nil,nil)
		return
	}
	infos,err := goods_dao.DB.GetSomeGoodsCollection(int64(limit),int64(offset),productids...)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,infos,nil)
	return
}
