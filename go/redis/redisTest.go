package main

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

var sql = ``

func main() {
	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {
		println(err)
		os.Exit(1)
	}
	var a = []byte{49, 0, 50}
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		c.Send("set", fmt.Sprintf("foo%d", i), string(a))
		reply, _ := c.Do("get", fmt.Sprintf("foo%d", i))
		fmt.Println(reply)
	}

	c.Send("set", "b", "2\x008")
	reply, _ := c.Do("get", "b")
	fmt.Println(reply)
	reply, _ = c.Do("get", "a")
	fmt.Println(reply)
}
