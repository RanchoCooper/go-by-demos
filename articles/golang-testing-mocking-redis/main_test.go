package main

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var redisServer *miniredis.Miniredis

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

func setup() {
	redisServer = mockRedis()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
}

func teardown() {
	redisServer.close()
}

func testDoRedisStuffDataExists(t *testing.T) {
	setup()
	defer teardown()

	redisClient.Set("data", "something here", time.Minute)

	result := doSomeRedisStuff()

	assert.True(t, result, "We expect the 'data' to be in Redis")
}

func testDoRedisStuffDataDoesNotExist(t *testing.T) {
	setup()
	defer teardown()

	result := doSomeRedisStuff()

	assert.False(t, result, "We expect the 'data' to not be in Redis")
}