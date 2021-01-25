package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	var (
		s1 string
		isOk bool
	)
	var regs = []string{
		`192\.168\.(\d+)\.(\d+)`,
		`10\.(\d+)\.(\d+)\.(\d+)`,
	}
	ip1 := "192.168.1.10"
	ip2 := "10.10.20.31"
	s1, isOk = BatchReg(regs, ip1)
	fmt.Println(s1, isOk)

	s1, isOk = BatchReg(regs, ip2)
	fmt.Println(s1, isOk)

	isOk = IsEmail("jasonminghao@163.com")
	fmt.Println(isOk)

	isOk = IsQq("7609772973")
	fmt.Println(isOk)

	isOk = IsValidMobile("15914727381")
	fmt.Println(isOk)

	// 敏感词
	 r := redis.Options{
		Addr:               "127.0.0.1:6379",
		DB:                 0,
	}
	conn := redis.NewClient(&r)
	isOk = IsSensitiveFilterRedis(conn, "caonima nima")
	fmt.Println(1, isOk)

	msg := "nidaye de"
	for _, val := range msg {
		isOk, _ := regexp.MatchString(string(val), `[a-Z0-9]`)
		fmt.Println(isOk)
		fmt.Println(string(val))
	}
}