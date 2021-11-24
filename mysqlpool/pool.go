package mysqlpool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"zht-go-server/config"
)

var MysqlDB = make(map[string]*sql.DB)

func init() {
	for name, cfg := range config.MysqlConfig {
		MysqlDB[name] = conn(cfg)
	}
}

func conn(cfg config.MsqlCfg) *sql.DB {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db
}

func Pool(name string) *sql.DB {
	if name == "" {
		name = "default"
	}
	rPool, ok := MysqlDB[name]
	if ok {
		return rPool
	}
	return nil
}
