package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type DelGoodsCartListData struct {
	Uid 		uint64 		`json:"uid"`
	ProductId  uint64 		`json:"product_id"`

}

func DelGoodsCartList (c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsCartListData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.ProductId<=0||Data.Uid<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = cart_dao.DB.DelGoodsCart(Data.Uid,Data.ProductId)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
	}
	common.BuildResp(c,nil,nil)
	return

}