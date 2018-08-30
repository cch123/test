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

	var ch = make(chan *nats.Msg, 10)
	sub, err := nc.QueueSubscribeSyncWithChan("tasks", "workers", ch)
	if err != nil {
		fmt.Println(err)
		return
	}

	for msg := range ch {
		fmt.Println(string(msg.Data), msg.Reply, msg.Sub, msg.Subject, err)
	}
	sub.Unsubscribe()

}
