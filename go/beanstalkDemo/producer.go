package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kr/beanstalk"
)

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		fmt.Println("beanstakd not started")
		os.Exit(1)
	}
	id, err := c.Put([]byte("hello"), 1, 0, 120*time.Second)
	fmt.Println(id, err)
}
