package user_handler

//
//import (
//	"LuXiStores/user/service"
//	"github.com/gin-gonic/gin"
//	"github.com/mojocn/base64Captcha"
//)
//
//func GenerateCaptcha(c *gin.Context) {
//	//session := sessions.Default(c)
//	captchaConfig := service.GetCaptchaConfig()
//	//create base64 encoding captcha
//	//创建base64图像验证码
//	config := captchaConfig.ConfigCharacter
//	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//	verify_id, digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id, config)
//	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
//	//session.Set("captchaId", captchaId)
//	c.JSON(200, gin.H{
//		"code": 0,
//		"msg":  "获取验证码成功",
//		"data": gin.H{
//			"base64png": base64Png,
//			"verify_id": verify_id,
//		},
//	})
//}
