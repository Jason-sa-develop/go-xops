package util

import (
	"testing"
)

func TestIp(t *testing.T) {
	var ips = []string{"192.168.1.100", "10.20.32.22"}
	sortIps(ips)

	getLocalIp()
}