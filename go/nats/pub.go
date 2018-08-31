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

	err = nc.Publish("tasks", []byte("start from here"))
	nc.Publish("tasks", []byte("aaa"))
	nc.Publish("tasks", []byte("bbb"))
	nc.Publish("tasks", []byte("ccc"))
	nc.Publish("tasks", []byte("ddd"))
	nc.Publish("tasks", []byte("eee"))
	nc.Publish("tasks", []byte("fff"))
	nc.Publish("tasks", []byte("ggg"))
	nc.Publish("tasks", []byte("hhh"))
	nc.Publish("tasks", []byte("end here"))
	// 要加了 flush 对面才能收到，看来内部也有缓冲区
	nc.Flush()
}
