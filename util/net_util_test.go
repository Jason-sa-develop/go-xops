package util

import (
	"fmt"
	"testing"
)

func TestNet(t *testing.T) {
	res := checkSocket("tcp", "172.16.1.21", 6379, 2)
	fmt.Println(res.M)
	fmt.Println(res.S)

}