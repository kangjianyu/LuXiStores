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
	MysqlClient *k_client.MysqlClient
)

func Init() {
	MysqlClient = newMysqlClient()
	RedisClient = newRedisClient()

}
func newMysqlClient() (mysqlClient *k_client.MysqlClient) {
	const prefix = "MysqlInit "
	dataSourceName := "root:KANG345876@qq.com@tcp(127.0.0.1:3306)/luxistores?charset=utf8"
	var err error
	mysql, err := gorm.Open("mysql", dataSourceName)
	mysqlClient = k_client.NewMysqlClient(mysql)
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
