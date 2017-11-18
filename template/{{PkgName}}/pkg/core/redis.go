package core

import (
	"fmt"

	environ "github.com/katsew/go-getenv"
	"github.com/go-redis/redis"
)

const (
	RedisHostKey     = "REDIS_HOST"
	RedisHostDefault = "127.0.0.1"
	RedisPortKey     = "REDIS_PORT"
	RedisPortDefault = "7379"
	RedisPassKey     = "REDIS_PASSWORD"
	RedisPassDefault = ""
	RedisDBKey       = "REDIS_DB"
	RedisDBDefault   = "0"
)

var redisClient *redis.Client

func init() {

	addr := fmt.Sprintf(
		"%s:%s",
		environ.GetEnv(RedisHostKey, RedisHostDefault),
		environ.GetEnv(RedisPortKey, RedisPortDefault),
	)
	opts := redis.Options{
		Addr:     addr,
		Password: environ.GetEnv(RedisPassKey, RedisPassDefault).String(),
		DB:       environ.GetEnv(RedisDBKey, RedisDBDefault).Int(),
	}
	redisClient = redis.NewClient(&opts)
}

func GetRedisInstance() *redis.Client {
	return redisClient
}