package order_handler

import (
	"LuXiStores/common"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
	"strings"
)


func CheckAliPay(c *gin.Context){
	urls:=c.Request.URL.RawQuery
	Sign := c.Query("sign")
	signtype := c.Query("sign_type")
	urls = strings.ReplaceAll(urls,"&sign="+Sign,"")
	urls = strings.ReplaceAll(urls,"&sign_type="+signtype,"")
	waitsign,_ := url.QueryUnescape(urls)
	signlist := strings.Split(waitsign,"&")
	sort.Strings(signlist)
	waitsign = strings.Join(signlist,"&")
	alipaypub := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAg9jXuSLH4dRUOWxfSLcROeGS9aSsxf9bsgLS5t3INJxCwYQ6jT/qIno3uvf57AxYSoagzrXTdczSSVgwp/WLBvqUDeWeioFVXqW7eb3/44sugkYhFsTuRWlzKP5mdzFgs7Get+6h+2WQIBxAVjyP7RacRWUSV0645LPI7PMrgsDvA1FRtxQTBi4mbxU+KuFJUG8bdnPIq0LgIpMum5ZOFDaVNfXIvwdqVEYPF06Q7Up9bRJSSJwtcXLNAAQvC1QhEhEsSwmueVrkJtyA1wtMQywVdRuBTJmtolAiilNhUNZSQFttJDQo4PxG6ttjAvH5Pbm3b75fIk90fdjCkiaKlwIDAQAB"
	sign ,err :=base64.StdEncoding.DecodeString(Sign)
	public, _ := base64.StdEncoding.DecodeString(alipaypub)
	PublicKey,err := x509.ParsePKIXPublicKey(public)

	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}

	err =rsa.VerifyPKCS1v15(PublicKey.(*rsa.PublicKey),crypto.SHA3_256,[]byte(sign),[]byte(waitsign))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}



}