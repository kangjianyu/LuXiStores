package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
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
	if err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	value,err := user_dao.Rds.GetUserToken(token)
	uid,err:=strconv.Atoi(value)
	if err!=nil||value==""{
		common.BuildResp(c,nil,common.ErrAuth)
		return
	}

	err = user_dao.DB.UpdateUserInfo(uint64(uid),updatedata.Password,updatedata.Email,updatedata.Phone)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)

}
