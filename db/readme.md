# 简介
针对go-redis进行二次封装，实现自定义功能


# 使用指南
## 连接redis
xredis支持两种方式初始化redis对象
方式一：通过配置文件（xx.ini）方式，支持多个redis服务端
方式二：手动指定连接信息，只返回一个redis连接服务端


### 1、连接多个redis
1.准备配置文件（路径：./config.ini）
```ini
;[redis对象名称_redis]  这里指定的对象名称，会映射到map的key中，
;addr=ip:port
;pool=连接池数量
;db=库

[devops_redis]
addr=127.0.0.1:6379
pool=30
db=0

[task_redis]
addr=127.0.0.1:6380
pool=30
db=0
```

2.指定配置文件并初始化连接对象
```go
if dsCli, err := newRedisFile("./config.ini"); err !=nil {
    panic("redis init error")
}
rdsCli.getKeys("devops", "dev-*")
```

3.使用不同的redis服务端连接对象
```go

```


### 2、连接一个redis
