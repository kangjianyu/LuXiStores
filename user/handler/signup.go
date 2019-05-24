package handler

import (
	"LuXiStores/user/service"
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"github.com/mojocn/base64Captcha"
)

//UserSignup 用户注册
func UserSignup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	verify_code := c.PostForm("verify_code")
	verify_id := c.PostForm("verify_id")
	if verify_code == "" || verify_id == "" {
		c.JSON(400, gin.H{
			"code": 1003,
			"msg":  "请输入验证码",
		})
		return
	} else {
		if !base64Captcha.VerifyCaptcha(verify_id, verify_code) {
			c.JSON(400, gin.H{
				"code": 1004,
				"msg":  "验证码错误",
			})
			return
		}
	}
	if valid := CheckUserName(username); valid != true {
		c.JSON(400, gin.H{
			"code": 1005,
			"msg":  "该用户已存在",
		})
		return
	}
	code, err := service.AddUserInfo(username, password, email, phone, gender)
	if err != nil {
		switch errcode := code; errcode {
		case 1001:
			c.JSON(400, gin.H{
				"err_code":  code,
				"error_msg": "参数错误",
			})

		case 1002:
			c.JSON(500, gin.H{
				"err_code":  code,
				"error_msg": "注册失败",
			})
		default:
			c.JSON(400, gin.H{
				"code": 10000,
				"msg":  "未知错误",
			})
		}
		return
	} else {
		c.JSON(201, gin.H{
			"code": 0,
			"msg":  "注册成功",
		})
		return
	}
}

//CheckUsername 用户名验证
func CheckUserName(name string) (valid bool) {
	prefix := "CheckUserName"
	ret, err := service.GetUserInfoByName(name)
	if err == nil || ret.UserName != "" {
		valid = false
		log.Error(prefix, "check username error:%s", name)
		return
	}
	valid = true
	log.Info(prefix, "check accept username:%s", name)
	return
}
