package common

import (
	"github.com/go-redis/redis"
)

// redis 公共模块

// 检测redis是否可访问
func checkService(cli *redis.Client) error {
	var (
		err error
	)

	// 基于redis模块测试
	_, err = cli.Ping().Result()
	return err
}

// 删除正则匹配的key
func DeleteManyKeys(cli *redis.Client, key string) error {
	var (
		cursor uint64 = 0  // 当前游标
		count  int64  = 10 // 搜索数量（次）
		err    error
	)
	if err = checkService(cli); err != nil {
		return err
	}

	pipe := cli.Pipeline()

	for {
		var keys []string // 临时

		// 搜索key
		keys, cursor, err = cli.Scan(cursor, key, count).Result()

		// 删除key
		cli.Unlink(keys...)

		if err != nil {
			return err
		}
		if cursor == 0 {
			break
		}
	}
	_, err = pipe.Exec()

	return err

}

// 删除指定key
func DeleteKeys(cli *redis.Client, keys []string) error {
	var (
		err error
	)
	// 删除指定key
	pipe := cli.Pipeline()
	pipe.Unlink(keys...)
	_, err = pipe.Exec()
	return err
}
