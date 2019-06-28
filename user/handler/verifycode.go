package user_handler


import (
	"LuXiStores/common"
	"github.com/gin-gonic/gin"
	log"github.com/jeanphorn/log4go"
	"github.com/mojocn/base64Captcha"
	"sync"
)
type CaptchaConfig struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}
var (
	captchaConfig *CaptchaConfig
	captchaConfigOnce sync.Once
)

func GenerateCaptcha(c *gin.Context) {
	//session := sessions.Default(c)
	captchaConfig := GetCaptchaConfig()
	//create base64 encoding captcha
	//创建base64图像验证码
	config := captchaConfig.ConfigCharacter
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	verify_id, digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	//session.Set("captchaId", captchaId)
	data := gin.H{"base64png": base64Png,}
	c.SetCookie("verifyid", verify_id, 0, "/", "", false, false)
	log.Info("grnerate captcha succeed id:%s",verify_id)
	common.BuildResp(c,data,nil)
}

func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:              "",
			CaptchaType:     "character",
			VerifyValue:     "",
			ConfigAudio:     base64Captcha.ConfigAudio{},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height:             100,
				Width:              240,
				Mode:               base64Captcha.CaptchaModeAlphabet,
				IsUseSimpleFont:    false,
				ComplexOfNoiseText: 0,
				ComplexOfNoiseDot:  0,
				IsShowHollowLine:   false,
				IsShowNoiseDot:     false,
				IsShowNoiseText:    false,
				IsShowSlimeLine:    false,
				IsShowSineLine:     false,
				CaptchaLen:         4,
			},
			ConfigDigit:     base64Captcha.ConfigDigit{
				Width:10,
				Height:10,
			},
		}
	})
	return captchaConfig
}