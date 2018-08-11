package main

import (
	"fmt"
	"os"

	"github.com/amyangfei/redlock-go/redlock"
)

func main() {
	locker, err := redlock.NewRedLock([]string{
		"tcp://127.0.0.1:6379",
		"tcp://127.0.0.1:6380",
		"tcp://127.0.0.1:6381",
	})

	if err != nil {
		fmt.Printlnt(err)
		os.Exit(1)
	}

	expirity, err := locker.Lock("resource_name", 200)

	err := locker.UnLock()
	fmt.Println(err, expirity)
}
