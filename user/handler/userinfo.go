package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	log "github.com/jeanphorn/log4go"

)

type UpdateData struct {
	Password string `json:"password"`
	Email string    `json:"email"`
	Phone string	`json:"phone"`
}
func UpdatePassword(c *gin.Context){
	data,err := ioutil.ReadAll(c.Request.Body)
	token,err:=c.Cookie("update_token")
	updatedata := UpdateData{}
	err  = json.Unmarshal(data,&updatedata)
	if err!=nil||updatedata.Email==""||updatedata.Phone==""||updatedata.Password==""{
		log.Warn("input data error %v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	value,err := user_dao.Rds.GetUserToken(token)
	uid,err:=strconv.Atoi(value)
	if err!=nil||value==""{
		log.Warn("token valid %s",token)
		common.BuildResp(c,nil,common.ErrAuth)
		return
	}

	err = user_dao.DB.UpdateUserInfo(uint64(uid),updatedata.Password,updatedata.Email,updatedata.Phone)
	if err!=nil{
		log.Warn("update userinfo error :%v uid:%d",err,uid)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info("update succeed uid%d",uid)
	common.BuildResp(c,nil,nil)

}
