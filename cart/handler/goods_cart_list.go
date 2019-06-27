package cart_handler

import (
	"LuXiStores/cart/dao"
	"LuXiStores/common"
	"github.com/gin-gonic/gin"
	"strconv"
)


type GetGoodsCartListData struct {
	Uid uint64 		`json:"uid"`
	Offset uint64 	`json:"start"`
	Limit uint64	`json:"count"`
}
func GetGoodsCartList(c *gin.Context){
	uidstr := c.Query("uid")
	offsetstr := c.Query("offset")
	limitstr := c.Query("limit")
	uid,err := strconv.Atoi(uidstr)
	offset,err := strconv.Atoi(offsetstr)
	limit ,err := strconv.Atoi(limitstr)
	if err!=nil||uid<=0||offset<0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	if limit==0{
		limit=20
	}

	infos,err := cart_dao.DB.GetGoodsCartList(uint64(uid),uint64(limit),uint64(offset))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	has_next := false
	if len(infos)>=limit{
		has_next = true
	}
	data := gin.H{
		"cart":infos,
		"has_next":has_next,
		"offset":offset+len(infos),
	}
	common.BuildResp(c,data,nil)
	return
}