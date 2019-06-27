package category_dao

import "LuXiStores/common"

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
	ret := common.RedisClient.Rds.Set(key,value,-1)
	return ret.Err()
}

