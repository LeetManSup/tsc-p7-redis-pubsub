package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Hello from publisher! Message #%d", i)
		if err := rdb.Publish(ctx, topic, message).Err(); err != nil {
			fmt.Printf("Publish error: %v\n", err)
		} else {
			fmt.Printf("Published: %s\n", message)
		}
		time.Sleep(2 * time.Second)
	}
}
