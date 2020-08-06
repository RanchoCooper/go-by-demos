package main

import "github.com/go-redis/redis"

// https://itnext.io/golang-testing-mocking-redis-b48d09386c70

var redisClient *redis.Client

func doSomeRedisStaff() bool {
	_, err := redisClient.Get("data").Result()

	return err != redis.Nil
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               "localhost:6579",
		Password:           "",
		DB:                 0,
	})
}

func main() {
	redisClient = newRedisClient()
	doSomeRedisStaff()
}