package k_client

import (
	"github.com/go-redis/redis"
	log "github.com/jeanphorn/log4go"
	"runtime"
	"strings"
	"time"
)

type RedisClient struct {
	// 包装一层, 方便打日志和统计数据
	Rds *redis.Client
}

func NewRedisClient(rds *redis.Client) *RedisClient {
	return &RedisClient{Rds: rds}
}

func doLog(ret redis.Cmder) {
	pc, _, _, _ := runtime.Caller(2)
	detail := runtime.FuncForPC(pc)
	funcPath := strings.Split(detail.Name(), "/")
	callerName := funcPath[len(funcPath)-1]
	err := ret.Err()
	if err == nil {
		log.Info("(REDIS %s)|%s|exist", callerName, ret.Args())
	} else if err == redis.Nil { // 如果是为空错误, 则非ERROR级别
		log.Info("(REDIS %s)|%s|not exist", callerName, ret.Args())
	} else {
		log.Error("(REDIS %s)|%s|err:%s", callerName, ret.Args(), err)
	}
}

func (r *RedisClient) Ping() *redis.StatusCmd {
	return r.Rds.Ping()
}

func (r *RedisClient) Del(keys ...string) *redis.IntCmd {
	ret := r.Rds.Del(keys...)
	doLog(ret)
	return ret
}

func (r *RedisClient) Get(key string) *redis.StringCmd {
	ret := r.Rds.Get(key)
	doLog(ret)
	return ret
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := r.Rds.Set(key, value, expiration)
	doLog(ret)
	return ret
}
func (r *RedisClient) Hash(key string,field string,value interface{} ) *redis.BoolCmd{
	ret := r.Rds.HSet(key,field,value)
	doLog(ret)
	return ret
}
func (r *RedisClient) HDel (key string,field ...string)*redis.IntCmd{
	ret := r.Rds.HDel(key,field...)
	doLog(ret)
	return ret
}
func (r *RedisClient) HMGet(key string,field ...string)*redis.SliceCmd{
	ret := r.Rds.HMGet(key,field...)
	doLog(ret)
	return ret
}

//func (r *RedisClient) Expire(key string, expiration time.Duration) *redis.BoolCmd {
//	return
//}
//func (r *RedisClient) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
//	return
//}
//func (r *RedisClient) TTL(key string) *redis.DurationCmd {
//	return
//}
//func (r *RedisClient) Type(key string) *redis.StatusCmd {
//	return
//}
//func (r *RedisClient) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
//	return
//}
//func (r *RedisClient) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
//	return
//}
//func (r *RedisClient) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
//	return
//}
//func (r *RedisClient) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
//	return
//}
//
//func (r *RedisClient) GetSet(key string, value interface{}) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) Incr(key string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) IncrBy(key string, value int64) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) MGet(keys ...string) *redis.SliceCmd {
//	return
//}
//func (r *RedisClient) MSet(pairs ...interface{}) *redis.StatusCmd {
//	return
//}
//
//func (r *RedisClient) HDel(key string, fields ...string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) HExists(key, field string) *redis.BoolCmd {
//	return
//}
//func (r *RedisClient) HGet(key, field string) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) HGetAll(key string) *redis.StringStringMapCmd {
//	return
//}
//func (r *RedisClient) HIncrBy(key, field string, incr int64) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) HKeys(key string) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) HLen(key string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) HMGet(key string, fields ...string) *redis.SliceCmd {
//	return
//}
//func (r *RedisClient) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
//	return
//}
//func (r *RedisClient) HSet(key, field string, value interface{}) *redis.BoolCmd {
//	return
//}
//func (r *RedisClient) HVals(key string) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) LLen(key string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) LPop(key string) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) LPush(key string, values ...interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) LPushX(key string, value interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) LRange(key string, start, stop int64) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) LRem(key string, count int64, value interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
//	return
//}
//func (r *RedisClient) LTrim(key string, start, stop int64) *redis.StatusCmd {
//	return
//}
//func (r *RedisClient) RPop(key string) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) RPopLPush(source, destination string) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) RPush(key string, values ...interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) RPushX(key string, value interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) SAdd(key string, members ...interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) SCard(key string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) SDiff(keys ...string) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) SDiffStore(destination string, keys ...string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) SInter(keys ...string) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) SInterStore(destination string, keys ...string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) SIsMember(key string, member interface{}) *redis.BoolCmd {
//	return
//}
//func (r *RedisClient) SMembers(key string) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) SMembersMap(key string) *redis.StringStructMapCmd {
//	return
//}
//func (r *RedisClient) SPop(key string) *redis.StringCmd {
//	return
//}
//func (r *RedisClient) SPopN(key string, count int64) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) ZAdd(key string, members ...Z) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZIncr(key string, member Z) *redis.FloatCmd {
//	return
//}
//func (r *RedisClient) ZCard(key string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZCount(key, min, max string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
//	return
//}
//func (r *RedisClient) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZRangeByScore(key string, opt ZRangeBy) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) ZRangeByScoreWithScores(key string, opt ZRangeBy) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZRank(key, member string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZRem(key string, members ...interface{}) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZRevRangeByScore(key string, opt ZRangeBy) *redis.StringSliceCmd {
//	return
//}
//func (r *RedisClient) ZRevRangeByScoreWithScores(key string, opt ZRangeBy) *redis.ZSliceCmd {
//	return
//}
//func (r *RedisClient) ZRevRank(key, member string) *redis.IntCmd {
//	return
//}
//func (r *RedisClient) ZScore(key, member string) *redis.FloatCmd {
//	return
//}
//func (r *RedisClient) Info(section ...string) *redis.StringCmd {
//	return
//}
