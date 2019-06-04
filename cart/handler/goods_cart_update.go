package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UpdateGoodsCartListData struct {
	Id 	uint64		`json:"id"`
	Uid uint64 		`json:"uid"`
	ProductId uint64 	`json:"product_id"`
	Quantity uint64	`json:"quantity"`
}


func UpdateGoodsCartList (c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := UpdateGoodsCartListData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Quantity<=0||Data.Uid<=0||Data.ProductId<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = cart_dao.DB.UpdateGoodsCart(Data.Uid,Data.ProductId,Data.Quantity)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}
