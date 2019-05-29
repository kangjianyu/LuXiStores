package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"time"
)


type ForgetData struct {
	Username string
	Email string
	Phone string
}
func ForgetPassword(c *gin.Context){
	data,err := ioutil.ReadAll(c.Request.Body)
	forgetdata := ForgetData{}
	err = json.Unmarshal(data,&forgetdata)
	if err!=nil||forgetdata.Username==""||forgetdata.Email==""||forgetdata.Phone==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	userinfo ,err :=user_dao.DB.GetUserInfoByUsername(forgetdata.Username)
	if err!=nil||userinfo.Uid==0{
		common.BuildResp(c,nil,common.ErrAuth)
		return
	}
	if userinfo.Phone!=forgetdata.Phone||userinfo.Email!=forgetdata.Email{
		common.BuildResp(c,nil,common.ErrAuth)
		return
	}
	token,err :=SetUpdateToken(userinfo.Username,userinfo.Uid)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	c.SetCookie("update_token", token, 0, "/", "", false, false)
	common.BuildResp(c,gin.H{"update_token": token,},nil)
	return

}

func SetUpdateToken(username string,uid uint64)(token string,err error){
	rand.Seed(time.Now().Unix())
	random := rand.Intn(1000)
	key := fmt.Sprintf("update_token_%s_%d",username,random)
	err = user_dao.Rds.SetUpdateToken(key,uid,1000*time.Second)
	if err!=nil{

	}
	return key,err
}
