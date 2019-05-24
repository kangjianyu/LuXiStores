package handler

import (
	"LuXiStores/user/service"
	"github.com/gin-gonic/gin"
)

func UserExit(c *gin.Context) {
	//prefix:= "UserExit"
	token, err := c.Cookie("sessionid")
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "failed",
			"message": "cookie无效",
		})
	}
	ok, err := service.DelUserTokenByRedis(token)
	if ok == 1 && err == nil {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "退出成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "failed",
			"message": "找不到该token",
		})
	}

}
