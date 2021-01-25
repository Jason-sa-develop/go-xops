package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gops/file"
	"io"
	"io/ioutil"
	"runtime"
	"strings"
)

// 获取cpu基础信息

func NumCpu() int {
	return runtime.NumCPU()
}

func CpuMHz() (mhz string, err error) {
	f := "/proc/cpuinfo"
	var (
		bs []byte
	)

	bs, err = ioutil.ReadFile(f)
	if err != nil {
		return
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))

	for {
		var lineBytes []byte
		lineBytes, err = file.ReadLine(reader)

		if err == io.EOF {
			return
		}

		line := string(lineBytes)
		if !strings.Contains(line, "MHz") {
			continue
		}

		arr := strings.Split(line, ":")
		if len(arr) != 2 {
			return "", fmt.Errorf("%s content format error", f)
		}

		// 去除两边空格
		return strings.TrimSpace(arr[1]), nil
	}
	return "", fmt.Errorf("no MHz in %s", f)
}

func main() {
	r1, rr1 := CpuMHz()
	fmt.Println(r1, rr1)
}
