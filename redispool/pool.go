package redispool

import (
	"github.com/gomodule/redigo/redis"
	"zht-go-server/config"
)

var RedisClient = make(map[string]*redis.Pool)

func init()  {
	for name,cfg := range config.RedisConfig{
		RedisClient[name] = conn(cfg)
	}
}

func conn(cfg config.RdsCfg)(*redis.Pool){
	return &redis.Pool{
		MaxIdle: cfg.MaxIdle,
		MaxActive: cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Wait: true,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr)
			if err != nil {
				return nil, err
			}
			if cfg.Auth != ""{
				if _, err := c.Do("AUTH", cfg.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cfg.Db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func Pool(name string) *redis.Pool{
	if name == ""{
		name = "default"
	}
	rPool, ok := RedisClient[name]
	if ok {
		return rPool
	}
	return nil
}

