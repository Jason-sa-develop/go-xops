package main

import (
	"fmt"
	"gops/file"
	"strconv"
	"strings"
)

type Loadavg struct {
	Avg1min        float64
	Avg5min        float64
	Avg15min       float64
	RunningProcess int64
	TotalProcess   int64
}


func LoadAvg()(*Loadavg, error){
	var (
		data string
		err error
	)

	loadAvg := Loadavg{}
	if data, err = file.ToTrimString("/proc/loadavg"); err != nil {
		return nil, err
	}

	// 以连续的空白字符为分隔符（返回列表）
	L := strings.Fields(data)

	if loadAvg.Avg1min, err = strconv.ParseFloat(L[0], 64); err  !=nil {
		return nil, err
	}
	if loadAvg.Avg5min, err = strconv.ParseFloat(L[1], 64); err  !=nil {
		return nil, err
	}
	if loadAvg.Avg15min, err = strconv.ParseFloat(L[2], 64); err  !=nil {
		return nil, err
	}
	// 按`/`符号切割 active process/sleep process
	processes := strings.SplitN(L[3], "/", 2)
	if len(processes) != 2 {
		return nil, err
	}

	if loadAvg.RunningProcess, err = strconv.ParseInt(processes[0], 10, 64); err  !=nil {
		return nil, err
	}
	if loadAvg.TotalProcess, err = strconv.ParseInt(processes[1], 10, 64); err  !=nil {
		return nil, err
	}

	return &loadAvg, nil

}

func main() {
	r, s := LoadAvg()
	fmt.Println(r,s)
}