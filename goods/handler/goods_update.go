package goods_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0||Data.Stock<0||Data.Name==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = goods_dao.DB.UpdateGoodsInfo(Data.Id,Data.CategoryId,Data.Name,Data.Subtitle,Data.MainImage,Data.SubImages,Data.Detail,Data.Price,Data.Stock)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}

func UpdateGoodsStatus(c *gin.Context){
	Indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsStatusData{}
	err = json.Unmarshal(Indata,&Data)
	if err!=nil||Data.Id<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	if Data.Status=="publish"{
		err = goods_dao.DB.UpdateGoodsStatus(Data.Id,1)
		if err!=nil{
			common.BuildResp(c,nil,common.ErrInternal)
			return
		}
		common.BuildResp(c,nil,nil)
		return
	}

	if Data.Status=="unpulish"{
		err = goods_dao.DB.UpdateGoodsStatus(Data.Id,0)
		if err!=nil{
			common.BuildResp(c,nil,common.ErrInternal)
			return
		}
		common.BuildResp(c,nil,nil)
		return
	}
	common.BuildResp(c,nil,common.ErrParam)
	return

}
