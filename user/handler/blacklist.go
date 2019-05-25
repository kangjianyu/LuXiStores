package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"io/ioutil"
	"strings"
)
type Body struct {
	Uids []string
}
func BlackListAdd(c *gin.Context){
	key := "blacklist"
	data,_ := ioutil.ReadAll(c.Request.Body)
	body := new(Body)
	err := json.Unmarshal(data,&body)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		log.Warn("uids==[]",body.Uids)
		return
	}

	for _,id := range body.Uids{
		err := user_dao.Rds.BlackListAdd(key,id)
		if err!=nil{
			common.BuildResp(c, nil, common.ErrInternal)
			return
		}
	}

	common.BuildResp(c, nil, nil)
	return
}
func BlackListDel(c *gin.Context){
	key := "blacklist"
	data,_ := ioutil.ReadAll(c.Request.Body)
	body := new(Body)
	err := json.Unmarshal(data,&body)
	if len(body.Uids)==0||err!=nil{
		log.Warn("uids==[]",body.Uids)
		common.BuildResp(c, nil, common.ErrParam)
		return
	}

	err = user_dao.Rds.BlackListDel(key,body.Uids...)
	fmt.Println(err)
	if err!=nil{
		common.BuildResp(c, nil, common.ErrInternal)
		return
	}

	common.BuildResp(c,nil,nil)
	return
}
func BlackListCheck(c *gin.Context){
	key := "blacklist"
	uids := c.Query("uids")
	uidlist := strings.Split(uids,",")
	if len(uidlist)==0{
		common.BuildResp(c,nil,common.ErrParam)
		log.Warn("uids==[]")
		return
	}

	value,err:=user_dao.Rds.BlackListCheck(key,uidlist...)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	data  := make(map[string]bool)
	for i,id:=range uidlist{
		if value[i]!="1"{
			data[id] = false
		}else{
			data[id] = true
		}
	}

	common.BuildResp(c,data,nil)
	return

}

