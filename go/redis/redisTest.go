package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	fmt.Printf("%#v\n", c)
	defer c.Close()
	if err != nil {
		println(err)
		os.Exit(1)
	}

	fmt.Println(time.Now())
	for i := 0; i < 40000; i++ {
		_, err := c.Do("sadd", "myset1", fmt.Sprint(rand.Intn(1000000)))
		_, err = c.Do("sadd", "myset3", fmt.Sprint(rand.Intn(1000000)))

		if err != nil {
			fmt.Println(err)
		}
	}

}
