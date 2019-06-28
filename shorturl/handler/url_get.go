package shorturl_handler

import (
	"LuXiStores/common"
	"LuXiStores/shorturl/dao"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
)

func GetLongUrl(c *gin.Context){
	prefix := "GetLongUrl"
	shorurl := c.Param("ww")
	if shorurl==""{
		log.Warn(prefix,"inpu data error",)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	info ,err := shorturl_dao.DB.GetShorUrl(shorurl)
	if err!=nil||info.LongUrl==""{
		log.Warn(prefix,"shorturl not exist:%s",shorurl)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	log.Info(prefix,"get succeed shorturl:%s",shorurl)
	c.Redirect(301,info.LongUrl)
	return

}
