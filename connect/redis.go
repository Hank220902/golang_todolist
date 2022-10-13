package connect

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)


func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
	return client
}

