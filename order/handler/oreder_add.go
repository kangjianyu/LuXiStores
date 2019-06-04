package order_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"LuXiStores/order/dao"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"io"
	"io/ioutil"
	"math/rand"
	"sort"
	"strings"
	"time"
)
type Alipay struct {
	App_id string
	charset string

}
type AddOrderData struct {
	UserId uint64 `json:"user_id"`
	ProductId uint64 `json:"product_id"`

}
func AddOrder(c *gin.Context){
	inData,err := ioutil.ReadAll(c.Request.Body)
	Data := &AddOrderData{}
	err = json.Unmarshal(inData,&Data)
	if err!=nil||Data.UserId<=0||Data.ProductId<=0{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}


	Uuid,err := uuid.GenerateUUID()
	goodsinfo ,err := goods_dao.DB.GetGoodInfoDetail(Data.ProductId)
	alipayKey := fmt.Sprintf("alipay_user=%d_product=%d_price=%.2f_%s",Data.UserId,Data.ProductId,goodsinfo.Price,Uuid)

	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	header,sign := alipaySign(goodsinfo.Name,Data.UserId,Data.ProductId,goodsinfo.Price)
	err = order_dao.Rds.SetOrderId(alipayKey,sign,time.Minute*15)
	err = order_dao.DB.AddOrder(alipayKey)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	header["trade_id"] = alipayKey
	common.BuildResp(c,header,nil)
	return


}


func alipaySign(title string,uid uint64,productId uint64,price float64) (map[string]interface{},string){
	header := make(map[string]interface{})
	header["app_id"] = "2019060265478144"
	header["charset"] = "utf-8"
	header["method"] = "alipay.trade.wap.pay"
	header["sign_type"] = "RSA2"
	header["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	header["version"] = "1.0"
	header["biz_content"] = make(map[string]string)
	header["biz_content"].(map[string]string)["productId"] = "1"
	header["biz_content"].(map[string]string)["userId"] = "2"
	header["biz_content"].(map[string]string)["price"] = "0.01"
	header["biz_content"].(map[string]string)["title"] = "我要买东西"
	signs := []string{}
	for i,x := range header{
		if value,ok := x.(string);ok==true{
			signs = append(signs,fmt.Sprintf(i+"="+value+"&"))
		} else{
			sign := ""
			sign += i+"={"
			for j,y := range x.(map[string]string){
				sign += fmt.Sprintf("\"%s\":"+"\"%v\",",j,y)
			}
			sign += "}&"
			signs = append(signs,sign)
		}
	}
	sort.Strings(signs)
	sign := ""
	sign = strings.Join(signs,sign)
	sign = sign[0:len(sign)-1]
	fmt.Println(sign)
	pivkey,err :=x509.ParsePKCS1PublicKey([]byte("MIIEpAIBAAKCAQEAtGCoAJWqZ+SG13048bSWTTxsphfz1nZg3Js3ovztdZqfw2aB74rfjC964WbhInLMTZ0HEXHy/oR6tciPFNGliptr32fOHCcLSdOGY2r1j0fT8ETpx6cskyoktUbcX4Muqgw4LUZU0BRfg+EsvWOmSe3WANxq8A0ECtvZGpr3itMNquuyxQY+6mA5LoQWbZC2ERYvzr/gbDVuWHlA6/WnsdHwVOwBsLfrod2396vWomfH+R75H0bAz5N7PqKsX0SaehK5oO90Z3lkBYSmec4UN1djBxUjC2llU0RxFIH8WwpnTtzfz1Z74wS90JJFVxVDZ/CvZUIoUsEAX1z3ZBbaLQIDAQABAoIBAQCL8CDmszZM+8KRE5lGC7A/o403HoYR8C0deV4kmM0w3BDua2yLBtZ/z6YpkMNBEobl/9kn85ttUiJRPZOjtzIS4plB7Sq+NJxRXkV4g9aWnkcStKQaPNwcICnyaVM21nMxgeFjXpkWBXhEvEEVfWXZHSdV66sbKT3lnsJEHc3GaMBUlYe+ld3aV+VrJcFqf99qFRCnhumC0xfe6zB+SFOuJ/mOHyjGmAaZHPrVis/gH5nXJT84Qv4pVbR4ISVQRXu3xTYX2A9kq6ZYtCZx4B/ZTKU3u2nGCxB5RSFhywuzEN2tGg/EdaS614kVrSE4NOjKgXYcAU0S5LewRns+4mgRAoGBAOhSOJ2nM4H7sqSNcnUiqAA6ST4AdQiif1G0ChTarnHrQ/oHUZnq5NONJvaoiLi2VvdJoC/GHa71o1dhtlo99zEcp4G0GgfJzE6Bsnvp8thwAJ1pcDj9Bf4EMYjpxk2I+sQlu4DwyQiMG5BAxhjBqLcd/Qm5InpreAhklG0qdDBTAoGBAMbDHDER/i2IwrqT6+zRA9l0+l3f0P2P2JP7Z1czNcJHtt40gUvZjssPRzhIpM/Dljun05WUruiAjqc6KdI71TdFNLCjokrl17m3AR3xrwmfcGXEeJQXWK0nIMpI1yBRtK5jjtPKKemZbpsoXgiq0njBiC1EkQ/xtWO+c/lueXt/AoGBAISSVshwH05vaRPJu6ToL+JhYGZHMIHK6Mig6pfX6nALhvDouEIS7p1iEPf0WIC/XIUkuIpKjanHdnxov/xjG+okpdm4ApqrJzEthcJ8UB3+W/t3rZh3mrHHhtTQQl8AackAly0POkjsWtZIgEKkUDienkSsJuag6RAxBRn+fesNAoGADfaO+HOHI7PD2k+h91UHrDMnk4ixqd59HIhAzkNut2NKWXney3FRMrq0CiQwT9gxqac0mgGD+Blv7BeN8JL7e5KFDROxxwk2inlsvqnH3ikxQDT5M44gUKm7B/ruAfs7cjTUR9Sf9SUuPAAt+vXlK06NPoDen/we/g7XHuK/7gkCgYBQ67RNCVzHvGwugsAOZYlvTN8QlEeK+h5UbSkUGjkUQ/BnP1Dc/ovLZ/CCgOApGtY1X2sDAiM/KNzPcn9cm7mEIRyZRphjpkiTWQbic4VIqjUfKGaquXMit4zuyKwPDgzskLI/nMAwdf2KCH/3HDvN750ICkizSH1JzIPvzDnTdg=="))
	h := sha256.New()
	h.Write([]byte(sign))
	h.Sum(nil)
	pk, err := ParsePrivateKey(pemPriKey)
	rsa.SignPKCS1v15(nil,)
	rsa.EncryptPKCS1v15(rand.Read(),pivkey,[]byte(sign))
	util
	encoding.ParsePKCS1PrivateKey(encoding.FormatPrivateKey(privateKey))

	s := base64.StdEncoding.EncodeToString(a)
	fmt.Println(s)
	header["sign"] = s

	return header,s
}

