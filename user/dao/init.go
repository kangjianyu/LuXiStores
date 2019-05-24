package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/jeanphorn/log4go"
	"time"
)

var (
	redisClient *redis.Client
	mysqlClient *sql.DB
)

func Init() {
	mysqlClient = newMysqlClient()
	redisClient = newRedisClient()

}
func newMysqlClient() (mysqlClient *sql.DB) {
	const prefix = "MysqlInit "
	dataSourceName := "root:KANG345876@qq.com@tcp(127.0.0.1:3306)/luxistores?charset=utf8"
	var err error
	mysqlClient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Error(prefix+"connect mysql error : dataSourceName: %s", dataSourceName)
		panic("MysqlInit failed")
	}
	log.Info(prefix+"connect mysql success : dataSourceName: %s", dataSourceName)
	return
}
func newRedisClient() (RedisClient *redis.Client) {
	if true {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			PoolSize: 100,
		})
		log.Info("set redis addr :%s", "127.0.0.1:6379")
	} else {
		RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "127.0.0.1:6379",
			SentinelAddrs: make([]string, 0),
			DialTimeout:   1 * time.Second,
			ReadTimeout:   1 * time.Second,
			WriteTimeout:  1 * time.Second,
			PoolSize:      100,
			PoolTimeout:   5 * time.Second,
		})
		log.Info("set redis addr: master: %s, sentinel:%v", "127.0.0.1:6379", "")
	}
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic("RedisInit error!")
	}
	log.Info("redis PING ==> PONG: %+v, error :%v", pong, err)
	log.Info("RedisInit connect redis success")

	return
}
