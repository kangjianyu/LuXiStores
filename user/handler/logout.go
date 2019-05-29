package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	token,err := c.Cookie("sessionid")
	if err != nil {
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = user_dao.Rds.DelUserToken(token)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrRedisKeyNotExist)
		return
	}
	common.BuildResp(c,nil,nil)
	return

}
