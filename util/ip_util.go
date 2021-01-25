package util

import (
	"fmt"
	"github.com/vearne/go-ping"
	"net"
	"sort"
	"strings"
	"time"
)

/*
	IP相关工具
*/

// 判断是否内网IP
func isIntranetIp(){

}


// ping检测 ip：指定目标IP地址; count：总次数; timeout：超时时间
func pingCmd(ip string, count int, timeout time.Duration) (res pingInfo) {
	var (
		pinger *ping.Pinger
		err    error

		setPkg  int
		recvPkg int
	)
	st := time.Now()
	pinger, err = ping.NewPinger(ip)
	et := time.Since(st)

	if err != nil {
		return
	}

	// 转换为秒
	timeout = timeout * time.Second

	pinger.Count = count
	pinger.Timeout = timeout

	pinger.OnFinish = func(stats *ping.Statistics) {
		setPkg = stats.PacketsSent
		recvPkg = stats.PacketsRecv
	}

	pinger.Run()

	if setPkg != recvPkg {
		// 请求失败
		res.M = fmt.Sprintf("ping超时%v秒未响应", timeout)
		return res
	}
	res.M = fmt.Sprintf("ping请求正常，耗时%v", et)
	res.S = true
	return res
}

// 对IP进行排序
func sortIps(ips []string) {
	sort.Strings(ips)
}

// 获取本机IP
func getLocalIp() (data []string) {
	var (
		addrs []net.Addr
		err   error
	)
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}

	for _, val := range addrs {
		match := isIpv4(val.String())
		if match {
			ipSplit := strings.Split(val.String(), "/")
			data = append(data, ipSplit[0])
		}

	}

	return data
}

// 判断是否为IPv4地址
func isIpv4(s string) (ret bool) {
	var ipv4Regs = []string{`\d+\.\d+\.\d+\.\d+`}
	_, ret = batchReg(ipv4Regs, s)
	return ret
}

