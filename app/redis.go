package app;

import (
	"context"
	"github.com/go-redis/redis/v8"
)

const REDIS_LOCATION_SUFFIXE = "/loc"
const REDIS_ADDRESS_LIST = "address"

func Store(key string, value string) error {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:" + Conf.REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
	})
	
	return rdb.Set(context.Background(), key, value, 0).Err()
}

func Get(key string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:" + Conf.REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
	})

	return rdb.Get(context.Background(), key).Result()
}

func AddToAddressList(address string) error{
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:" + Conf.REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
	})

	return rdb.LPush(context.Background(), REDIS_ADDRESS_LIST, address).Err()
}

func GetAddressList() ([]string, error){
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:" + Conf.REDIS_PORT,
        Password: "", // no password set
        DB:       0,  // use default DB
	})

	return rdb.LRange(context.Background(), REDIS_ADDRESS_LIST, 0, -1).Result()
}