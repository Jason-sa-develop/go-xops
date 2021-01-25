package util

import (
	"fmt"
	"net"
	"time"
)

/*
	网络相关工具
*/

type pingInfo struct {
	S bool   // 状态
	M string // 描述信息
}

// 端口检测 protocol：指定协议；ip：指定目标IP地址；port：指定目标端口
func checkSocket(protocol string, ip string, port string, timeout time.Duration) (res pingInfo) {
	if protocol == "" {
		return res
	}
	timeout = timeout * time.Second

	// ip:port
	addr := fmt.Sprintf("%v:%v", ip, port)

	// 发送检测请求
	st := time.Now()
	_, err := net.DialTimeout(protocol, addr, timeout)
	et := time.Since(st)

	if err != nil {
		// 请求失败
		res.M = fmt.Sprintf("ping超时%v秒未响应", timeout)
		return res
	}
	// 请求成功
	res.M = fmt.Sprintf("ping请求正常，耗时%v", et)
	res.S = true
	return res

}

