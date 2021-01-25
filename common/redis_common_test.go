package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	rdsCli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
	})
	err := DeleteManyKeys(rdsCli, "test-*")
	fmt.Println(err)
}
