package config

import (
	"log"

	redis "github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // default DB
	})

	// test connection by sending ping and printing the output
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis PING Says: ", pong)
	return client
}
