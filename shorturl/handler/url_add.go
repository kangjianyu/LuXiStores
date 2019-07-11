package shorturl_handler

import (
	"LuXiStores/common"
	"LuXiStores/shorturl/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"io/ioutil"
	"strings"
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
	urlstr :=transform(urlid)
	shorturl := fmt.Sprintf("kjy%s",urlstr)
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


func transform(urlid int64) string{
	arr := []string{"0","1","2","3","4","5","6","7","8","9",
		"a","b","c","d","e","f","g","h","i","j","k","l","n",
		"m","o","p","q","r","s","t","u","v","w","x","y","z",
		"A","B","C","D","E","F","G","H","I","J","K","L","N",
		"M","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	res := []string{}
	for true{
		s := urlid/62
		y := urlid%62
		res = append(res,arr[y])
		if s==0{
			break
		}
		urlid = s
	}
	fmt.Println(res)
	l := len(res)-1
	for i:=0;i<l/2;i++{
		res[i],res[l-i] = res[l-i],res[i]
	}
	return strings.Join(res,"")
}

func getNextUrl() (int64,error){
	ret := common.RedisClient.IncrBy("short/u",1)
	return ret.Val(),ret.Err()
}
