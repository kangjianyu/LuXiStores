package order_dao

import (
	"LuXiStores/common"
	"time"
)

var Rds iRedis = redisImpl{}

type iRedis interface {
	SetOrderId(key string,values string,TTl time.Duration)(error)

}
type redisImpl struct{}

func (redisImpl) SetOrderId(key string,values string,TTl time.Duration) (error) {
	return  common.RedisClient.Set(key,values,TTl).Err()
}


