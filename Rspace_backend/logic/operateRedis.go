// 封装redis的写读操作

package logic

import (
	"Rspace_backend/dao"
	"context"
	"time"
)

var ctx = context.Background()

// redis插入值
func RedisSet(key, val string, time time.Duration) error {
	if err := dao.RDB.Set(ctx, key, val, time).Err(); err != nil { //在redis中存mykey只存time秒钟，time秒后自动消除
		//do something
		// panic(err)
		return err
		// return err
	}
	return nil
}

// redis读取值
func RedisGet(key string) (string, error) {
	val, err := dao.RDB.Get(ctx, key).Result()
	if err != nil {
		// panic(err)
		return "", err
	}
	return val, nil
}
