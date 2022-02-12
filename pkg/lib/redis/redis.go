package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

//OptionsFields 配置项字段
type OptionsFields struct {
	Host     string //主机地址
	Port     int    //端口
	Password string //密码
	Database int    //库序号
}

var redisInstance *redis.Client

//InitRedis 初始化redis连接
func InitRedis(opts OptionsFields) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		Password: opts.Password,
		DB:       opts.Database,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	redisInstance = rdb
}

//GetInstance 获取redis连接实例
func GetInstance() *redis.Client {
	return redisInstance
}
