package user_handler

import (
	"LuXiStores/common"
	user_dao "LuXiStores/user/dao"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	log "github.com/jeanphorn/log4go"
	"github.com/mojocn/base64Captcha"
	"strconv"
	"time"
)

//UserLogin 用户登录
func Login(c *gin.Context) {
	//prefix := "UserLogin"
	username := c.PostForm("username")
	password := c.PostForm("password")
	verify_code := c.PostForm("verify_code")
	verifyId,err := c.Cookie("verifyid")
	// 入参错误
	if verify_code == "" || verifyId == "" || username == "" || password == "" {
		log.Warn(`verify_code == "" || verify_id == "" || username == "" || password == ""`)
		common.BuildResp(c, nil, common.ErrParam)
		return
	}
	// 校验码不正确
	if !base64Captcha.VerifyCaptcha(verifyId, verify_code) {
		log.Info("VerifyCaptcha error:%s",verify_code)
		common.BuildResp(c, nil, common.ErrCaptcha)
		return
	}

	userinfo, err := user_dao.DB.GetUserInfoByUsername(username)
	// 用户不存在
	if err != nil {
		log.Warn("get userinfo id%s,error:%v,",username,err)
		common.BuildResp(c, nil, common.ErrAuth)
		return
	}

	// 密码错误
	if password != userinfo.Password {
		log.Info("password error:%d",password)
		common.BuildResp(c, nil, common.ErrAuth)
		return
	}

	token,  err := SetToken(userinfo.Uid)
	if err != nil{
		log.Error("set redis error:%v,uid:%d",err,userinfo.Uid)
		return
	}
	c.SetCookie("sessionid", token, 0, "/", "", false, false)
	log.Info("login succeed %s",username)
	common.BuildResp(c, gin.H{"token": token}, nil)
	return
}

//SetCookie 设置用户cookie
func SetToken(userid uint64) (token string,  err error) {
	prefix := "SetCookie"
	Uuid, err := uuid.GenerateUUID()
	if err != nil {
		log.Error(prefix, "get uuid error %v", err)
		return
	}
	struid := strconv.FormatUint(userid, 10)
	err = user_dao.Rds.SetUserToken(Uuid, struid, time.Hour*24*7)
	if err != nil {
		return
	}
	token = Uuid
	log.Info("set cookie in redis %s",Uuid)
	return
}

//CheckCookie cookie验证
func CheckCookie(uuid string) (value string, err error) {
	prefix := "CheckCookie"
	value, err = user_dao.Rds.GetUserToken(uuid)
	if err != nil {
		log.Error(prefix, "check usercookie error:%v", err)
		return
	}
	log.Info(prefix, "check usercookie:%s success ", uuid)
	return
}