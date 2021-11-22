package main

import (
    "fmt"
    "time"

    //    "github.com/gin-gonic/gin"
    "zht-go-server/redispool"

    //    "github.com/gorilla/websocket"
)

func main() {
    vtime := time.Now().UnixNano()/1e6
    rc := redispool.Select("default")

    //配置分布式锁
    lock :=rc.Lock("lock:zeng","zeng",5)
    //获取分布式锁
    ok,owner,err := lock.Set()

    fmt.Println("lock",ok,err)
    if ok{
        res,err := rc.HMSet("zeng:map:001","t1","001","t2","002","t3","003")
        fmt.Println("hmset",res,err)

        rmap,err :=rc.HGetAll("zeng:map:001")
        fmt.Println("hmget",rmap,err)

        fmt.Printf("[SUM]:%v ms\n" , time.Now().UnixNano()/1e6 - vtime)
    }
    //释放分布式锁
    lock.Release(owner)
}
