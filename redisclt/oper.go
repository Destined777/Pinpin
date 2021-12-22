package redisclt

import (
	"Pinpin/global"
	"github.com/go-redis/redis"
	"time"
)

func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return global.RedisClient.Set(key, value, expiration)
}
func Get(key string) *redis.StringCmd {
	return global.RedisClient.Get(key)
}
func Del(keys string) *redis.IntCmd {
	return global.RedisClient.Del(keys)
}
