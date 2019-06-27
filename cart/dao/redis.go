package cart_dao

var Rd iRedis = redisimpl{}

type iRedis interface {
	//SetCartInfo(uid int64,info string)

}

type redisimpl struct {

}