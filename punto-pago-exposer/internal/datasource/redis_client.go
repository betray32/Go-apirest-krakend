package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func redis_manager() {
	// Set up Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Password, if any
		DB:       0,  // Database number (default is 0)
	})

	// Set a value in Redis
	err := client.Set("nombre", "Juan", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get a value from Redis
	name, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)
}
