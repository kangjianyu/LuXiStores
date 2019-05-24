package user_handler

import (
	"LuXiStores/common"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	//prefix:= "Logout"

	common.BuildHttpResp(c, gin.H{"greet": "hello"}, nil)
	return

}
