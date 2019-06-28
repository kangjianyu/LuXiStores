package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	log"github.com/jeanphorn/log4go"

)

type UpdateGoodsCartListData struct {
	Id 	uint64		`json:"id"`
	Uid uint64 		`json:"uid"`
	ProductId uint64 	`json:"product_id"`
	Quantity uint64	`json:"quantity"`
}


func UpdateGoodsCartList (c *gin.Context){
	prefix := "UpdateGoodsCartList"
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsCartListData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Quantity<=0||Data.Uid<=0||Data.ProductId<=0{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = cart_dao.DB.UpdateGoodsCart(Data.Uid,Data.ProductId,Data.Quantity)
	if err!=nil{
		log.Warn(prefix,"update error:%v,",err,)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed Data:%v",Data)
	common.BuildResp(c,nil,nil)
	return
}
