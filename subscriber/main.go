package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	topic := "notifications"
	sub := rdb.Subscribe(ctx, topic)

	ch := sub.Channel()

	fmt.Println("Subscribed to", topic)

	for msg := range ch {
		fmt.Printf("Received message: %s\n", msg.Payload)
	}
}
