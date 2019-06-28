package receiver_handler

import (
	"LuXiStores/common"
	"LuXiStores/receiver/dao"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"strconv"
)

func GetGoodsReceiverAddressList(c *gin.Context){
	prefix := "GetGoodsReceiverAddressList"
	strid := c.Query("uid")
	uid,err := strconv.Atoi(strid)
	if err!=nil||uid<=0{
		log.Warn(prefix,"input data error%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	infos,err:=receiver_dao.DB.GetGoodsReceiverAddress(uint64(uid))
	if err!=nil{
		log.Warn(prefix,"get default address error:%v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info(prefix,"succeed data:%v",infos)
	common.BuildResp(c,infos,nil)
	return
}

