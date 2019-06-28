package category_dao

import (
	"LuXiStores/common"
	"time"
)

var Rds iRedis = redisImpl{}
type iRedis interface {
	SetCategoryInfo(key string,value interface{}) (error)
	GetCategoryInfo(key string) (string,error)
}
type redisImpl struct{}

func (redisImpl) GetCategoryInfo(key string) (string, error) {
	ret := common.RedisClient.Rds.Get(key)
	return ret.Val(),ret.Err()
}

func (redisImpl) SetCategoryInfo(key string,value interface{}) (error) {
	ret := common.RedisClient.Rds.Set(key,value,time.Second*5*60)
	return ret.Err()
}

