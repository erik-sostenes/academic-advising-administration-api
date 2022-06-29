package repository

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	syncRedis sync.Once
	redisConnection *redis.Client
	redisPingErr error
)

// LoadRedisConnection create the connection to redis
func LoadRedisConnection() (*redis.Client, error) {
	syncRedis.Do(func() {
		redisConnection = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprint(os.Getenv("REDIS_HOST") +":"+ os.Getenv("REDIS_PORT")),
			Password: "",
			DB: 0,
		})

		statusCmd :=  redisConnection.Ping(context.TODO())
		redisPingErr = statusCmd.Err()
	})
	return redisConnection, redisPingErr
}

// NewRedis create the redis instance 
func NewRedis() *redis.Client {
	rdb, err := LoadRedisConnection()
	if err != nil {
		panic(err)
	}
	return rdb
}
