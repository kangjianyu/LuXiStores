package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"io/ioutil"
	"time"
)
type SignUp struct {
	Username string
	Password string
	Email string
	VerifyCode string
	Phone string
}

func UserSignUp(c *gin.Context){
	data,err := ioutil.ReadAll(c.Request.Body)
	verifyId,err := c.Cookie("verifyid")
	signupData := SignUp{}
	err = json.Unmarshal(data,&signupData)
	//入参错误
	if err!=nil||signupData.VerifyCode==""||signupData.Username==""||signupData.Phone==""||signupData.Email==""||signupData.Password==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	//验证码错误
	if !base64Captcha.VerifyCaptcha(verifyId, signupData.VerifyCode) {
		common.BuildResp(c, nil, common.ErrCaptcha)
		return
	}
	//用户名验证
	if !CheckUserName(signupData.Username){
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	info := user_dao.UserInfo{
		Username:   signupData.Username,
		Password:   signupData.Password,
		Email:      signupData.Email,
		Status:     0,
		Phone:      signupData.Phone,
		CreateTime:time.Now(),
		UpdateTime:time.Now(),
	}
	if err := user_dao.DB.AddUserInfo(info);err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}


//CheckUsername 用户名验证
func CheckUserName(name string) (valid bool) {
	userinfo,err :=user_dao.DB.GetUserInfoByUsername(name)
	if err!=nil{
		return
	}
	if userinfo.Username==""{
		return true
	}
	return false
}
