package main

import (
	"flag"
	"time"

	"github.com/garyburd/redigo/redis"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
)

func main() {
	flag.Parse()
	pool = newPool(*redisServer)
	conn := pool.Get()
	defer conn.Close()
}
