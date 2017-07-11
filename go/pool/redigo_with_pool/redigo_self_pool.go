package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	pool := &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			/*
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
				if _, err := c.Do("SELECT", db); err != nil {
					c.Close()
					return nil, err
				}
			*/
			return c, nil
		},
	}
	defer pool.Close()

	c := pool.Get()
	b, err := c.Do("info")
	if err != nil {
		fmt.Println("errr", err)
		return
	}

	defer c.Close()
	fmt.Printf("%s\n", b)
}
