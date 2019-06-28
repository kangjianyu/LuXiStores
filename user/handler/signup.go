package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"github.com/mojocn/base64Captcha"
	"io/ioutil"
	"math/rand"
	"time"
)
type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string 	`json:"email"`
	VerifyCode string `json:"verify_code"`
	Phone string `json:"phone"`
}

func UserSignUp(c *gin.Context){
	data,err := ioutil.ReadAll(c.Request.Body)
	verifyId,err := c.Cookie("verifyid")
	signupData := SignUp{}
	err = json.Unmarshal(data,&signupData)
	//入参错误
	if err!=nil||signupData.VerifyCode==""||signupData.Username==""||signupData.Phone==""||signupData.Email==""||signupData.Password==""{
		log.Info("input data error%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	//验证码错误
	if !base64Captcha.VerifyCaptcha(verifyId, signupData.VerifyCode) {
		log.Info("VerifyCaptcha error%s",signupData.VerifyCode)
		common.BuildResp(c, nil, common.ErrCaptcha)
		return
	}
	//用户名验证
	if !CheckUserName(signupData.Username){
		log.Info("username has exist %s",signupData.Username)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	uid,err := getNextUid()
	info := user_dao.UserInfo{
		Uid:uint64(uid),
		Username:   signupData.Username,
		Password:   signupData.Password,
		Email:      signupData.Email,
		Status:     0,
		Phone:      signupData.Phone,
	}
	if err := user_dao.DB.AddUserInfo(info);err!=nil{
		log.Warn("signup error %v",err)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info("signup succeed name%s",signupData.Username)
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

func getNextUid() (int64,error){
	rand.Seed(time.Now().Unix())
	value := rand.Intn(100)
	ret := common.RedisClient.Rds.IncrBy("max_uid",int64(value))
	err := ret.Err()
	if ret.Val()==int64(value){
		uid,err := user_dao.DB.GetMaxUid()
		if uid==0{
			ret = common.RedisClient.Rds.IncrBy("max_uid",100000)
		} else{
			ret = common.RedisClient.Rds.IncrBy("max_uid",int64(uid))
		}
		return ret.Val(),err

	}
	return ret.Val(),err
}