package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	client.Do(ctx, "set", "k1", "n1")
	res, err := client.Do(ctx, "get", "k1").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(res)
	}
}
