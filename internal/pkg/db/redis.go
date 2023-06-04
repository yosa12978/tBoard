package db

import (
	"os"
	"sync"

	"github.com/go-redis/redis"
)

var (
	rdb       *redis.Client
	redisOnce sync.Once
)

func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASS"),
			DB:       0,
		})
		if err := rdb.Ping().Err(); err != nil {
			panic(err)
		}
	})
	return rdb
}
