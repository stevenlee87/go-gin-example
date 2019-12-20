package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stevenlee87/go-gin-example/pkg/setting"
)

func main() {
	var RedisConn *redis.Pool

	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
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
	}

	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "key4", "rabbit4")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}
