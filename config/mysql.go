package config

import "time"

type MsqlCfg struct {
	//地址 ex: 127.0.0.1:6379
	DSN string
	//最小连接数
	MaxIdleConns int
	//最大连接数
	MaxOpenConns int
	//超时时间
	ConnMaxLifetime time.Duration
}

var MysqlConfig = map[string]MsqlCfg{
	"default": MsqlCfg{
		DSN:             "root:123456@tcp(192.168.123.84:3306)/test",
		MaxIdleConns:    10,
		MaxOpenConns:    50,
		ConnMaxLifetime: 1 * time.Minute,
	},
}
