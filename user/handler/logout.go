package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
)

func Logout(c *gin.Context) {
	token,err := c.Cookie("sessionid")
	if err != nil ||token==""{
		log.Info("exit error not find %v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = user_dao.Rds.DelUserToken(token)
	if err!=nil{
		log.Warn("del session error %v",err)
		common.BuildResp(c,nil,common.ErrRedisKeyNotExist)
		return
	}
	log.Info("exit succeed %s",token)
	common.BuildResp(c,nil,nil)
	return

}
