package user_handler

import "github.com/gin-gonic/gin"

//
//import (
//	"LuXiStores/user/service"
//	"github.com/gin-gonic/gin"
//	"github.com/hashicorp/go-uuid"
//	log "github.com/jeanphorn/log4go"
//	"github.com/mojocn/base64Captcha"
//	"math/rand"
//	"strconv"
//	"time"
//)
//
////UserLogin 用户登录
//func Login(c *gin.Context) {
//	prefix := "UserLogin"
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	verify_code := c.PostForm("verify_code")
//	verify_id := c.PostForm("verify_id")
//	if verify_code == "" && verify_id == "" {
//		c.JSON(401, gin.H{
//			"code": 1004,
//			"msg":  "请输入验证码",
//		})
//		return
//	} else {
//		if !base64Captcha.VerifyCaptcha(verify_id, verify_code) {
//			c.JSON(401, gin.H{
//				"code": 1005,
//				"msg":  "验证码错误",
//			})
//			return
//		}
//	}
//	if username != "" && password != "" {
//		userinfo, err := service.GetUserInfoByName(username)
//		if err != nil {
//			log.Error(prefix, "get user:%s error:%v ", username, err)
//			c.JSON(401, gin.H{
//				"code": 1003,
//				"msg":  "该用户不存在",
//			})
//			return
//		} else if password == userinfo.PassWord {
//			token, valid, err := SetToken(userinfo.UId)
//			if err != nil || valid != "OK" {
//
//			}
//			c.SetCookie("sessionid", token, 0, "/", "", false, false)
//			c.JSON(200, gin.H{
//				"err_code":  0,
//				"error_msg": "登录成功",
//				"data": gin.H{
//					"token": token,
//				},
//			})
//			return
//		} else {
//			c.JSON(401, gin.H{
//				"err_code":  1001,
//				"error_msg": "用户名或密码错误",
//			})
//			return
//		}
//	} else {
//		c.JSON(401, gin.H{
//			"err_code":  1002,
//			"error_msg": "请输入用户名密码",
//		})
//		return
//	}
//
//}
//
////SetCookie 设置用户cookie
//func SetToken(userid uint64) (token string, valid string, err error) {
//	prefix := "SetCookie"
//	rand.Seed(time.Now().Unix())
//	random := rand.Intn(1000)
//	Uuid, err := uuid.GenerateUUID()
//	if err != nil {
//		log.Error(prefix, "get uuid error %v", err)
//		return
//	}
//	struid := strconv.FormatUint(userid, 10)
//	valid, err = service.SetUserTokenByRedis(Uuid, struid, 604800+int64(random))
//	token = Uuid
//	if err != nil {
//		log.Error(prefix, "set cookie error:%v ", err)
//		return
//	}
//	log.Info("set cookie in redis %s")
//	return
//}
//
////CheckCookie cookie验证
//func CheckCookie(uuid string) (value string, err error) {
//	prefix := "CheckCookie"
//	value, err = service.GetUserTokenByRedis(uuid)
//	if err != nil {
//		log.Error(prefix, "check usercookie error:%v", err)
//		return
//	}
//	log.Info(prefix, "check usercookie:%s success ", uuid)
//	return
//}

////UserLogin 用户登录
func Login(c *gin.Context) {
	//prefix := "UserLogin"
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	verify_code := c.PostForm("verify_code")
	verify_id := c.PostForm("verify_id")
	if verify_code == "" && verify_id == "" {
		c.JSON(401, gin.H{
			"code": 1004,
			"msg":  "请输入验证码",
		})
		return
	}
}
