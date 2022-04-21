package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/limitcool/blog/global"
	"time"
)

// redigo 功能封装
//设置 RedisConn 为 redis.Pool（连接池）并配置了它的一些参数：
var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", global.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if global.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", global.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     global.RedisSetting.MaxIdle,
		MaxActive:   global.RedisSetting.MaxActive,
		IdleTimeout: global.RedisSetting.IdleTimeout,
	}
	return nil
}

// Set a key/value time:(expire)过期时间
func Set(key string, data interface{}, time int) error {
	// RedisConn.Get方法从redis连接池中拿到redis连接
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return nil
	}
	return nil
}

// exists check a key 检查key是否在redis中
func Exists(key string) bool {
	// RedisConn.Get方法从redis连接池中拿到redis连接
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

// 通过key从redis获取值,值为[]byte类型
func Get(key string) ([]byte, error) {
	// RedisConn.Get方法从redis连接池中拿到redis连接
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, err
}

// 通过key 从redis中删除key
func Delete(key string) (bool, error) {
	// RedisConn.Get方法从redis连接池中拿到redis连接
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

// 通过key模糊搜索进行批量删除key
func LikeDeletes(key string) error {
	// RedisConn.Get方法从redis连接池中拿到redis连接
	conn := RedisConn.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return nil
	}
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
