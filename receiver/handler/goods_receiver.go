package receiver_handler

import (
	"LuXiStores/common"
	"LuXiStores/receiver/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetGoodsReceiverAddressList(c *gin.Context){
	strid := c.Query("uid")
	uid,err := strconv.Atoi(strid)
	if err!=nil||uid<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	infos,err:=receiver_dao.DB.GetGoodsReceiverAddress(uint64(uid))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	common.BuildResp(c,infos,nil)
	return
}

