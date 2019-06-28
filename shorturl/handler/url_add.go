package shorturl_handler

import (
	"LuXiStores/common"
	"LuXiStores/shorturl/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
)
type AddShortUrlData struct {
	LongUrl string `json:"long_url"`

}

func AddShortUrl(c *gin.Context){
	prefix := "AddShortUrl"
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddShortUrlData{}
	err = json.Unmarshal(indata,&Data)
	if Data.LongUrl==""||err!=nil{
		log.Warn(prefix,"input data error:%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	urlid,err := getNextUrl()
	shorturl := fmt.Sprintf("kjy%d",urlid)
	err = shorturl_dao.DB.AddShortUrl(shorturl,Data.LongUrl)
	if err!=nil{
		log.Warn(prefix,"insert error:%v,url:%s",err,Data.LongUrl)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	data := fmt.Sprintf("/m78/%s",shorturl)
	log.Info(prefix,"add succeed url:%s",data)
	common.BuildResp(c,data,nil)
	return

}


func getNextUrl() (int64,error){
	ret := common.RedisClient.IncrBy("short/u",1)
	return ret.Val(),ret.Err()
}