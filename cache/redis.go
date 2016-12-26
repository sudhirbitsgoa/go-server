package cache

import (
	"fmt"

	"gopkg.in/redis.v5"
)

var Client *redis.Client

/**
 * redis client
 */
func RedisConnect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}
