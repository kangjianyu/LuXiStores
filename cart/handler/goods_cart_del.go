package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	log"github.com/jeanphorn/log4go"

)

type DelGoodsCartListData struct {
	Uid 		uint64 		`json:"uid"`
	ProductId  uint64 		`json:"product_id"`

}

func DelGoodsCartList (c *gin.Context){
	prefix := "DelGoodsCartList"
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := DelGoodsCartListData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.ProductId<=0||Data.Uid<=0{
		log.Warn(prefix,"input data error:%v",err)

		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = cart_dao.DB.DelGoodsCart(Data.Uid,Data.ProductId)
	if err!=nil{
		log.Warn(prefix,"del error:%v,uid:%d,proId:%d",err,Data.Uid,Data.ProductId)
		common.BuildResp(c,nil,common.ErrInternal)
	}
	log.Info(prefix,"succeed uid:%d,proId:%d",Data.Uid,Data.ProductId)
	common.BuildResp(c,nil,nil)
	return

}