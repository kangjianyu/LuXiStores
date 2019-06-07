package shorturl_handler

import (
	"LuXiStores/common"
	"LuXiStores/shorturl/dao"
	"github.com/gin-gonic/gin"
)

func GetLongUrl(c *gin.Context){
	shorurl := c.Param("ww")
	if shorurl==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	info ,err := shorturl_dao.DB.GetShorUrl(shorurl)
	if err!=nil||info.LongUrl==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	c.Redirect(301,info.LongUrl)
	return

}
