package main

import (
	"fmt"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = nc.Publish("tasks", []byte("oh yeah yea"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("aaa"))
	// 要加了 flush 对面才能收到，看来内部也有缓冲区
	nc.Flush()
}
