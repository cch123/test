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
		fmt.Println("beanstalkd not started")
		os.Exit(1)
	}

	for {
		id, body, err := c.Reserve(5 * time.Hour)
		fmt.Println(id, string(body), err)
	}
}
