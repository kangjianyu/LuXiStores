package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)

type GoodsInfoData struct {
	CategoryId	uint64 		`json:"category_id"`
	Name		string		`json:"name"`
	Subtitle	string		`json:"subtitle"`
	MainImage	string		`json:"main_image"`
	SubImages	string		`json:"sub_images"`
	Detail		string		`json:"detail"`
	Price		float64		`json:"price"`
	Stock		uint64		`json:"stock"`
	Status		int64		`json:"status"`
}


func AddGoodsInfo(c *gin.Context){
	prefix := "AddGoodsInfo"
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := GoodsInfoData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.CategoryId<=0||Data.Price<0||Data.Stock<0{
		log.Warn(prefix,"input data error:v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.AddGoodsInfo(Data.CategoryId,Data.Name,Data.Subtitle,Data.MainImage,Data.SubImages,Data.Detail,Data.Price,Data.Stock,Data.Status)
	if err!=nil{
		log.Warn(prefix,"add goodsinfo error:%v,Data:%v",err,Data)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"add goodsinfo succeed data:%v",Data)
	common.BuildResp(c,nil,nil)
	return
}

