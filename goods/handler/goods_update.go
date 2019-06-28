package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)
type UpdateGoodsData struct {
	Id          uint64      `json:"id"`
	CategoryId	uint64 		`json:"category_id"`
	Name		string		`json:"name"`
	Subtitle	string		`json:"subtitle"`
	MainImage	string		`json:"main_image"`
	SubImages	string		`json:"sub_images"`
	Detail		string		`json:"detail"`
	Price		float64		`json:"price"`
	Stock		uint64		`json:"stock"`
}
type UpdateGoodsStatusData struct {
	Id          uint64      `json:"id"`
	Status 		string 		`json:"status"`
}
func UpdateGoodsInfo(c *gin.Context){
	prefix := "UpdateGoodsInfo"
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0||Data.Stock<0||Data.Name==""{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.UpdateGoodsInfo(Data.Id,Data.CategoryId,Data.Name,Data.Subtitle,Data.MainImage,Data.SubImages,Data.Detail,Data.Price,Data.Stock)
	if err!=nil{
		log.Warn(prefix,"update error:%v,goodsid:%d",err,Data.Id)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed id:%d",Data.Id)
	common.BuildResp(c,nil,nil)
	return
}

func UpdateGoodsStatus(c *gin.Context){
	prefix := ""
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsStatusData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	if Data.Status=="publish"{
		err = goods_dao.DB.UpdateGoodsStatus(Data.Id,1)
		if err!=nil{
			log.Warn(prefix,"update error:%v,goodsid:%d",err,Data.Id)
			common.BuildResp(c,nil,common.ErrInternal)
			return
		}
		log.Info(prefix,"succeed goodsid:%d is publish",Data.Id)
		common.BuildResp(c,nil,nil)
		return
	}

	if Data.Status=="unpulish"{
		err = goods_dao.DB.UpdateGoodsStatus(Data.Id,0)
		if err!=nil{
			log.Warn(prefix,"update error:%v,goodsid:%d",err,Data.Id)
			common.BuildResp(c,nil,common.ErrInternal)
			return
		}
		log.Info(prefix,"succeed goodsid:%d is unpublish",Data.Id)
		common.BuildResp(c,nil,nil)
		return
	}
	log.Warn(prefix,"input data error:%v",err)
	common.BuildResp(c,nil,common.ErrParam)
	return

}
