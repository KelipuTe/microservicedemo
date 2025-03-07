package ioc

import "github.com/redis/go-redis/v9"

func InitRedis() redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
