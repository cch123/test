package main

import (
	"fmt"
	"sync/atomic"

	"github.com/go-redis/redis"
)

type MyConn struct {
	c         *redis.Client
	opt       redis.Options
	available int32
}

func (c *MyConn) keepAliveCheck() {
	for {
		//time.Sleep(time.Second)
		_, err := c.c.Ping().Result()
		//fmt.Println(res, err)
		if err != nil {
			atomic.StoreInt32(&c.available, 1)
			continue
		}
		atomic.StoreInt32(&c.available, 0)
		atomic.LoadInt32(&c.available)
	}
}

func (c *MyConn) fakeGet() {
	for {
		c.c.Get("abc").Result()
	}
}

func getClient(addr string, password string, db int) *MyConn {
	c := &MyConn{}
	c.opt = redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	c.c = redis.NewClient(&c.opt)
	//for i := 0; i < 100; i++ {
	go c.keepAliveCheck()
	go c.fakeGet()
	//	}
	return c
}

func main() {

	conn := getClient("localhost:6379", "", 0)
	client := conn.c
	for {
		//time.Sleep(3 * time.Second)
		//fmt.Printf("%#v\n", conn)
	}

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}
