package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
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
	prefix:="AddGoodsCollection"
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddGoodsCollectionData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.ProductId<=0||Data.Uid<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.AddGoodsCollection(Data.Uid,Data.ProductId)
	if err!=nil{
		log.Warn(prefix,"add goodscollection error:%v,uid:%d,proId:%d",err,Data.Uid,Data.ProductId)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed uid:%d,productId:%d",Data.Uid,Data.ProductId)
	common.BuildResp(c,nil,nil)
	return
}

func DelGoodsCollection(c *gin.Context){
	prefix:="DelGoodsCollection"
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsCollectionData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Uid<=0||Data.ProductId<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.DelGoodsCollection(Data.Uid,Data.ProductId)
	if err!=nil{
		log.Warn(prefix,"del error%:v,data:%v",err,Data)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed data:%v",Data)
	common.BuildResp(c,nil,nil)
	return
}

func GetGoodsCollectionList(c *gin.Context){
	prefix := "GetGoodsCollectionList"
	uidstr := c.Query("uid")
	offsetstr := c.Query("offset")
	limitstr := c.Query("limit")
	uid,err := strconv.Atoi(uidstr)
	offset,err := strconv.Atoi(offsetstr)
	limit,err := strconv.Atoi(limitstr)

	if err!=nil||uid<=0||limit<0||offset<0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	if limit==0{
		limit = 20
	}
	infos,err := goods_dao.DB.GetGoodsCollectionByUid(int64(uid),int64(offset),int64(limit))
	if err!=nil{
		log.Warn(prefix,"get collection error:%v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed uid:%d",uid)
	common.BuildResp(c,infos,nil)
	return
}
