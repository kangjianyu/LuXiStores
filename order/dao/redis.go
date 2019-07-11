package order_dao

import (
	"LuXiStores/common"
	"time"
)

var Rds iRedis = redisImpl{}

type iRedis interface {
	SetOrderId(key string,values string,TTl time.Duration)(error)
	DecreaseStock(productId string,decrement int64) (int64,error)
	GetMaxUid(key string,value int64)(int64,error)
	SetStock(productId string,amount int64)(int64,error)
	CheckProductId(key string) (int64,error)
	GetOrderId(key string) (string,error)
	DelOrderId(key string)error
}
type redisImpl struct{}

func (redisImpl) DelOrderId(key string) error {
	ret := common.RedisClient.Del(key)
	return ret.Err()
}

func (redisImpl) GetOrderId(key string) (string, error) {
	ret := common.RedisClient.Get(key)
	return ret.Val(),ret.Err()
}

func (redisImpl) CheckProductId(key string) (int64,error) {
	ret := common.RedisClient.Exist(key)
	return ret.Val(),ret.Err()
}

func (redisImpl) SetStock(productId string, amount int64) (int64,error) {
	ret := common.RedisClient.IncrBy(productId,amount)
	err := ret.Err()
	err = common.RedisClient.Expire(productId,time.Second*3).Err()
	return ret.Val(),err
}

func (redisImpl) GetMaxUid(key string,value int64) (int64, error) {
	ret := common.RedisClient.DecrBy(key,value)
	return ret.Val(),ret.Err()
}

func (redisImpl) DecreaseStock(productId string, decrement int64) (int64,error) {
	ret := common.RedisClient.DecrBy(productId,decrement)
	err := ret.Err()
	err = common.RedisClient.Expire(productId,time.Second*3).Err()
	return ret.Val(),err
}

func (redisImpl) SetOrderId(key string,values string,TTl time.Duration) (error) {
	return  common.RedisClient.Set(key,values,TTl).Err()
}



