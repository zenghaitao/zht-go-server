package config

import "time"

type RdsCfg struct {
	//地址 ex: 127.0.0.1:6379
	Addr string
	//账号
	Auth string
	//db index
	Db int
	//最小连接数
	MaxIdle int
	//最大连接数
	MaxActive int
	//超时时间
	IdleTimeout time.Duration
}

var RedisConfig = map[string] RdsCfg{
	"default": RdsCfg{
		Addr:"192.168.123.84:6379",
		Auth:"",
		Db:0,
		MaxIdle: 10,
		MaxActive: 50,
		IdleTimeout: 30 * time.Second,
	},
	"t1": RdsCfg{
		Addr:"192.168.123.84:6379",
		Auth:"",
		Db:1,
		MaxIdle: 10,
		MaxActive: 50,
		IdleTimeout: 30 * time.Second,
	},
	"t2": RdsCfg{
		Addr:"192.168.123.84:6379",
		Auth:"",
		Db:2,
		MaxIdle: 10,
		MaxActive: 50,
		IdleTimeout: 30 * time.Second,
	},
}