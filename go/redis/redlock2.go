package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/redsync.v1"
)

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func newPools(servers []string) []redsync.Pool {
	pools := []redsync.Pool{}
	for _, server := range servers {
		pool := newPool(server)
		pools = append(pools, pool)
	}

	return pools
}

func main() {
	pools := newPools([]string{"127.0.0.1:6379"})
	rs := redsync.New(pools)
	m := rs.NewMutex("/lock")

	err := m.Lock()
	if err != nil {
		panic(err)
	}
	fmt.Println("lock success")
	unlockRes := m.Unlock()
	fmt.Println("unlock result: ", unlockRes)

}
