package util

import (
	"github.com/go-redis/redis"
	"regexp"
)

var (
	mobilePattern string = `^1([356789][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`
)

// 批量解析文本
func BatchReg(regs []string, str string) (string, bool) {
	return batchReg(regs, str)
}

func batchReg(regs []string, str string) (string, bool) {
	var ret bool // 是否匹配成功
	for _, reg := range regs {
		ret, _ = regexp.MatchString(reg, str)
		if ret {
			return str, ret
		}
	}
	return str, ret
}

// 是否邮箱
func IsEmail(s string) (ret bool) {
	ret, _ = regexp.MatchString(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]`, s)
	return ret
}

// 是否QQ
func IsQq(s string) (ret bool) {
	ret, _ = regexp.MatchString(`^\d{6,18}$`, s)
	return ret
}

// 验证合法手机号
func IsValidMobile(mobile string) (ret bool) {
	mobileLen := len(mobile)
	if mobileLen != 11 {
		return ret
	}
	ret, _ = regexp.MatchString(mobilePattern, mobile)
	return ret
}


/*
	敏感词过滤，支持对单一的消息做敏感词判断过滤，返回bool值，如果返回true 说明msg是敏感词，如果返回false 则不是敏感词
	1、词库存在redis中，通过从redis获取词库来做判断
	2、词库存在程序本地缓存中，通过本地缓存的词库来做判断
 */

// conn：redis连接对象
// msg：消息/内容
func IsSensitiveFilterRedis(conn *redis.Client, msg string) (ret bool) {
	// redis中的词库
	var sensitiveRepo = [2]string{"sensitive_01", "sensitive_02"}

	// todo 消息切割


	// 循环判断
	for _, s := range sensitiveRepo {
		res := conn.LRange(s, 0, -1)
		list, err := res.Result()
		if err != nil {
			return ret
		}
		return IsContainStr(msg, list)
	}
	return ret
}

// cache：本地缓存的词库
// msg：消息/内容
func IsSensitiveFilterCache(cache []string, msg string) (ret bool) {
	return IsContainStr(msg, cache)
}


