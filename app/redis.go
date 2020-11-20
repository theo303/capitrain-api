package app;

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func Store(key string, value string) error {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:" + Conf.REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
	})
	
	return rdb.Set(context.Background(), key, value, 0).Err()
}