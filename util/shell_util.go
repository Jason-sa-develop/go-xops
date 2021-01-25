package util

import (
	"os/exec"
	"strings"
)

// 执行shell命令
func RunCMD(cmd string, t string) (string, error) {
	// cmd：执行命令、t：shell类型
	switch t {
	case "bash":
		return bashShell(cmd)
	case "sh":
		return shShell(cmd)
	}
	return "", nil
}

// shell解释器由bash执行的结果
func bashShell(cmd string) (string, error) {
	return commandShell("/bin/bash", cmd)
}

// shell解释器由sh执行的结果
func shShell(cmd string) (string, error) {
	return commandShell("/bin/sh", cmd)
}

func commandShell(s string, cmd string) (string, error) {
	var (
		result *exec.Cmd
		data   string
		out    []byte
		err    error
	)
	result = exec.Command(s, "-c", cmd)

	// 读取数据
	if out, err = result.Output(); err != nil {
		// todo 记日志
		return data, err
	}

	// 将数据转换为string类型
	return strings.Trim(string(out), "\n"), err
}
