package common

import (
	"LuXiStores/k_client"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/jeanphorn/log4go"
	"github.com/jinzhu/gorm"
)

var (
	RedisClient *k_client.RedisClient
	MysqlClient *gorm.DB
)

func Init() {
	MysqlClient = newMysqlClient()
	RedisClient = newRedisClient()

}
func newMysqlClient() (mysqlClient *gorm.DB) {
	const prefix = "MysqlInit "
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/luxistores?charset=utf8&parseTime=true&loc=Local"
	var err error
	mysqlClient, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Error(prefix+"connect mysql error : dataSourceName: %s", dataSourceName)
		panic("MysqlInit failed")
	}
	log.Info(prefix+"connect mysql success : dataSourceName: %s", dataSourceName)
	return
}
func newRedisClient() (redisClient *k_client.RedisClient) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		PoolSize: 100,
	})

	redisClient = k_client.NewRedisClient(rds)
	log.Info("set redis addr :%s", "127.0.0.1:6379")

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic("RedisInit error!")
	}
	log.Info("redis PING ==> PONG: %+v, error :%v", pong, err)
	log.Info("RedisInit connect redis success")

	return
}
