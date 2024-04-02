// 连接mysql
package dao

import (
	"github.com/redis/go-redis/v9"
)

// var RDB =

var (
	RDB *redis.Client
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //redis默认端口6379
		Password: "123456",         // 密码
		DB:       0,                // use default DB
	})
}
