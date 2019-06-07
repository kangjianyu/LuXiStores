package shorturl_handler

import (
	"LuXiStores/common"
	"LuXiStores/shorturl/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)
type AddShortUrlData struct {
	LongUrl string `json:"long_url"`

}

func AddShortUrl(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddShortUrlData{}
	err = json.Unmarshal(indata,&Data)
	if Data.LongUrl==""||err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	urlid,err := getNextUrl()
	shorturl := fmt.Sprintf("kjy%d",urlid)
	err = shorturl_dao.DB.AddShortUrl(shorturl,Data.LongUrl)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	data := fmt.Sprintf("http://127.0.0.1:8080/m78/%s",shorturl)
	common.BuildResp(c,data,nil)
	return

}


func getNextUrl() (int64,error){
	ret := common.RedisClient.IncrBy("short/u",1)
	return ret.Val(),ret.Err()
}