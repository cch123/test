package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	opts := agent.Options{
		Addr: ":8080",
	}
	if err := agent.Listen(&opts); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		time.Sleep(time.Minute)
		go func() { time.Sleep(time.Minute) }()
	}

	time.Sleep(time.Hour)
}
