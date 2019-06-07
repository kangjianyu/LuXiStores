package order_handler

import (
	"LuXiStores/common"
	"LuXiStores/goods/dao"
	"LuXiStores/order/dao"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	random "math/rand"
	"sort"
	"strconv"
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
	if err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	alipayKey := fmt.Sprintf("alipay_user=%d_product=%d_price=%.2f_%s",Data.UserId,Data.ProductId,goodsinfo.Price,Uuid)
	header,sign,err := alipaySign(goodsinfo.Name,Data.UserId,Data.ProductId,goodsinfo.Price)
	err = order_dao.Rds.SetOrderId(alipayKey,sign,time.Minute*15)
	orderId,err := getNextOrderId()
	err = order_dao.DB.AddOrder(alipayKey,uint64(orderId),Data.ProductId)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	header["trade_id"] = alipayKey
	common.BuildResp(c,header,nil)
	return

}


func alipaySign(title string,uid uint64,productId uint64,price float64) (header map[string]interface{},sign string,err error){
	header = make(map[string]interface{})
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
	sign = strings.Join(signs,sign)
	sign = sign[0:len(sign)-1]
	fmt.Println(sign)
	privateKey,err := ioutil.ReadFile("/Users/kjy/Downloads/RSA签名验签工具_MAC_V3/RSA密钥/应用私钥_tmp.txt")
	h := sha256.New()
	h.Write([]byte(sign))
	var block *pem.Block
	block, _ = pem.Decode([]byte(privateKey))
	if block==nil{
		err = errors.New("私钥有错")
		return
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	d := h.Sum(nil)
	bs,err := rsa.SignPKCS1v15(rand.Reader,pk, crypto.SHA256, d)
	s := base64.StdEncoding.EncodeToString(bs)
	header["sign"] = s

	return header,s,err
}

func getNextOrderId()(int64,error){

	times := time.Now().Format("20060102150405")
	ret := common.RedisClient.IncrBy(times,1)
	fmt.Println(times)
	err := ret.Err()
	if ret.Val()==1{
		random.Seed(time.Now().Unix())
		ret =common.RedisClient.IncrBy(times,int64(random.Intn(1000)))
		err = common.RedisClient.Expire(times,5*time.Second).Err()
	}
	Id,err := strconv.Atoi(times+"0000")
	return int64(Id)+ret.Val(),err
}
func formatKey(raw, prefix, suffix string) (result []byte) {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var ll = 64
	var sl = len(raw)
	var c = sl / ll
	if sl%ll > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * ll
		var e = b + ll
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}