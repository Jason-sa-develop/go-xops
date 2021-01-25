package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/unknwon/goconfig"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type xRedis struct {
	RDS map[string]*redis.Client // 存放redis连接对象
}

// 初始化redis连接对象（配置文件）
func newRedisFile(filePath string) (*xRedis, error) {
	var (
		cfg             *goconfig.ConfigFile
		namesPrefix     []string                 // redis前缀
		sectionList     []string                 // section字段
		redisConnection map[string]*redis.Client // redis连接对象
		err             error
	)

	// 解析配置文件
	if cfg, err = goconfig.LoadConfigFile(filePath); err != nil {
		return nil, err
	}
	if sectionList = cfg.GetSectionList(); sectionList == nil {
		return nil, fmt.Errorf("未指定redis相关的section")
	}

	// 校验section
	for _, s := range sectionList {
		var isOK bool
		if isOK, err = regexp.MatchString(`[\w+]_redis`, s); err != nil {
			return nil, err
		}
		if isOK {
			split := strings.Split(s, "_")
			namesPrefix = append(namesPrefix, split[0])
		}

	}

	// 获取配置
	redisConnection = make(map[string]*redis.Client, len(namesPrefix))

	// 生成并保存redis连接对象
	for _, s := range namesPrefix {
		var (
			rdsCfg map[string]string // 存放配置
			addr   string
			pool   int
			db     int
		)
		sv := s + "_redis"

		if rdsCfg, err = cfg.GetSection(sv); err != nil {
			return nil, err
		}

		addr = rdsCfg["addr"]
		_pool := rdsCfg["pool"]
		_db := rdsCfg["db"]
		isDB := reflect.TypeOf(_db)
		if isDB.Name() == "string" {
			if db, err = strconv.Atoi(_db); err != nil {
				return nil, err
			}
		}
		isPool := reflect.TypeOf(_pool)
		if isPool.Name() == "string" {
			if pool, err = strconv.Atoi(_pool); err != nil {
				return nil, err
			}
		}

		obj := redis.NewClient(&redis.Options{
			Addr:     addr,
			DB:       db,
			PoolSize: pool,
		})
		if obj != nil {
			redisConnection[s] = obj
		}
	}

	return &xRedis{RDS: redisConnection}, nil
}

// 初始化redis连接对象（基于参数）
// 只返回一个连接对象
func newRedis(addr string, db int) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
	return conn, nil
}

// 连接指定对象
func (r *xRedis) Connect(name string) *redis.Client {
	return r.connect(name)
}

// 删除正则匹配key，只支持一个key，该key中的正则表达式必须是redis中可支持的
// 基于pipeline和延迟删除实现（性能好）
func (r *xRedis) DeleteManyKey(name string, key string) (err error) {
	var (
		keys   []string       // 临时存放key
		cursor uint64         // 游标
		count  int64    = 100 // 每次迭代的key数
	)

	cli := r.connect(name)
	pipeline := cli.Pipeline()

	for {
		if keys, cursor, err = cli.Scan(cursor, key, count).Result(); err != nil {
			return err
		}
		pipeline.Unlink(keys...)
		if cursor == 0 {
			break
		}
	}
	_, err = pipeline.Exec()
	return err
}

// 删除数组中的所有key
// 基于pipeline和延迟删除实现
func (r *xRedis) DeleteKeys(name string, keys []string) (err error) {
	cli := r.connect(name)

	pipeline := cli.Pipeline()
	for _, key := range keys {
		pipeline.Unlink(key)
	}
	_, err = pipeline.Exec()

	return err
}

// 统计指定多个key的数量
func (r *xRedis) LenKeys(name string, keys ...string) (result map[string]int) {
	result = make(map[string]int, len(keys)) // 初始化map

	for _, key := range keys {
		val := r.getKeys(name, key)
		result[key] = len(val)
	}
	return result
}

// 正则匹配key，返回匹配的所有key（基于scan）
func (r *xRedis) getKeys(name string, key string) (result []string) {
	var (
		cursor uint64 = 0 // 游标
		err    error
	)
	cli := r.connect(name)
	for {
		var keys []string
		if keys, cursor, err = cli.Scan(cursor, key, 10).Result(); err != nil {
			panic(fmt.Sprintf("【%v】搜索key错误", name))
		}
		result = append(result, keys...)
		if cursor == 0 {
			break
		}
	}
	return result

}

// 获取连接对象
func (r *xRedis) connect(name string) *redis.Client {
	obj, ok := r.RDS[name]
	if !ok {
		panic(fmt.Sprintf("【%v】连接对象不存在", name))
	}

	return obj
}
