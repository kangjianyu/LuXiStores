package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


type GetGoodsCartListData struct {
	Uid uint64 		`json:"uid"`
	Start uint64 	`json:"start"`
	Count uint64	`json:"count"`
}
func GetGoodsCartList(c *gin.Context){
	inData,err :=ioutil.ReadAll(c.Request.Body)
	Data := GetGoodsCartListData{}
	err = json.Unmarshal(inData,&Data)
	if err!=nil||Data.Uid<=0||Data.Start<0||Data.Count<0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	if Data.Count==0{
		Data.Count=20
	}

	infos,err := cart_dao.DB.GetGoodsCartList(Data.Uid,Data.Count,Data.Start)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	has_next := false
	if uint64(len(infos))>=Data.Count{
		has_next = true
	}
	data := gin.H{
		"cart":infos,
		"has_next":has_next,
		"offset":Data.Start+uint64(len(infos)),
	}
	common.BuildResp(c,data,nil)
	return
}