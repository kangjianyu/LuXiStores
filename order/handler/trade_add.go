package order_handler

import (
	"LuXiStores/common"
	"LuXiStores/order/dao"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"strings"
)



type AddTradeData struct {
	TradeId string `json:"trade_id"`
	Sign string `json:"sign"`
}

func AddTrade(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddTradeData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.TradeId==""||Data.Sign==""{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	check,err := order_dao.Rds.GetOrderId(Data.TradeId)
	if check!=Data.Sign{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	items := strings.Split(Data.TradeId,"_")
	user,product,price := "","",""
	amount,order,receiver := "","",""
	for _,x :=range items{
		if strings.Contains(x,"user"){
			index := strings.Index(x,"=")
			user = x[index+1:]
			fmt.Println(user,"id")
		}
		if strings.Contains(x,"product"){
			index := strings.Index(x,"=")
			product = x[index+1:]
			fmt.Println(product,"product")
		}
		if strings.Contains(x,"price"){
			index := strings.Index(x,"=")
			price = x[index+1:]
			fmt.Println(price,"price")
		}
		if strings.Contains(x,"amount"){
			index := strings.Index(x,"=")
			amount = x[index+1:]
			fmt.Println(amount,"amount")
		}
		if strings.Contains(x,"order"){
			index := strings.Index(x,"=")
			order = x[index+1:]
			fmt.Println(order,"order")
		}
		if strings.Contains(x,"receiver"){
			index := strings.Index(x,"=")
			receiver = x[index+1:]
			fmt.Println(receiver,"receiver")
		}

	}
	orderId,err := strconv.Atoi(order)
	userId,err := strconv.Atoi(user)
	productId,err := strconv.Atoi(product)
	receiverId,err := strconv.Atoi(receiver)
	prices,err := strconv.ParseFloat(price,10)
	amounts,err := strconv.ParseFloat(amount,10)
	amountPrice := prices*amounts
	err = order_dao.DB.AddTrade(int64(orderId),Data.TradeId,int64(userId),int64(productId),int64(receiverId),amountPrice,int64(amounts))
	err = order_dao.DB.OrderPayFinish(int64(orderId))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	err = order_dao.Rds.DelOrderId(Data.TradeId)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrRedisKeyNotExist)
		return
	}
	kafmsg := fmt.Sprintf("user=%s&price=%s&count=%s",user,price,amount)
	saramaProducer(kafmsg)

	common.BuildResp(c,nil,nil)
	return
}

func saramaProducer(value string)  {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V2_2_0_0

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer([]string{"127.0.0.1:9092",}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()


	// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
	msg := &sarama.ProducerMessage{
		Topic: "test",
	}
	//将字符串转化为字节数组
	msg.Value = sarama.ByteEncoder([]byte(value))
	//使用通道发送
	producer.Input() <- msg
}