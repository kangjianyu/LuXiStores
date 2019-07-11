package kafka

import (
	"LuXiStores/user/dao"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	log"github.com/jeanphorn/log4go"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

func InitConsumer(){
	ctx := context.TODO()
	go ClusterConsumer(ctx)
	return
}
func ClusterConsumer(ctx context.Context)  {
	//defer wg.Done()
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V2_2_0_0
	// init consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("consumer created")
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	partitionConsumer,err := consumer.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Warn("close consumer error,v",err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	ConsumerLoop:
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				vals := string(msg.Value)
				val := strings.Split(vals,"&")
				if len(val)<3{
					log.Warn("ErrParam,%s",val)
					continue
				}
				userstr := val[0][strings.Index(val[0],"=")+1:]
				userId,_ := strconv.Atoi(userstr)
				pricestr := val[1][strings.Index(val[1],"=")+1:]
				priceId,_ := strconv.ParseFloat(pricestr,10)
				countstr := val[2][strings.Index(val[1],"=")+1:]
				countId,err := strconv.ParseFloat(countstr,10)
				err = AddExp(int64(userId),priceId*countId)
				if err!=nil{
					log.Warn("rise exp error ")
					continue
				}
				log.Info("rise exp userid:%d,exp:%d",userId,priceId*100)
				consumed++
			case <-signals:
				break ConsumerLoop
			}
	}

}

func AddExp(useId int64,price float64) error{
	err := user_dao.PDB.AddUserExp(useId,price*100)
	return err
}







