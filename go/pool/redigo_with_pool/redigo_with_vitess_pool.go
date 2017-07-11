package main

import (
	"fmt"
	"log"
	"time"

	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
)

// ResourceConn : vitess 的 Resource 是一个 interface，这个 interface 只有 Close 一个接口
// 因为 redis.Conn 的 Close 定义和 vitess.Resource 的 Close 定义不一样
// 所以需要 wrap 一次，使 redis.Conn 转为 vitess 的 pool.Resource
// go 的通用连接池大概都是这么实现了
// TODO 不过通用连接池的探活应该怎么做呢？
type ResourceConn struct {
	redis.Conn
}

// Close : 实现 vitess 的 pool.Resource() 的 Close 方法
func (r ResourceConn) Close() {
	r.Conn.Close()
}

func main() {
	p := pools.NewResourcePool(
		func() (pools.Resource, error) {
			c, err := redis.Dial("tcp", ":6379")
			return ResourceConn{c}, err
		}, // factory 方法，这里是指获得一条 Conn 的方法
		1,           // capacity
		2,           // maxCap
		time.Minute, // idleTimeout
	)
	defer p.Close()
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	n, err := c.Do("INFO")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.StatsJSON())
	log.Printf("info=%s", n)
}
