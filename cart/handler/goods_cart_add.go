package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AddGoodsCartListData struct {
	Uid uint64 		`json:"uid"`
	ProductId uint64 	`json:"product_id"`
	Quantity uint64	`json:"quantity"`
}


func AddGoodsCartList(c	*gin.Context){
	indate,err := ioutil.ReadAll(c.Request.Body)
	Data := AddGoodsCartListData{}
	err = json.Unmarshal(indate,&Data)
	if err!=nil||Data.Uid<=0||Data.Quantity<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = cart_dao.DB.AddGoodsCart(Data.Uid,Data.ProductId,Data.Quantity)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}
