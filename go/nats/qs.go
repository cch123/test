package main

import (
	"fmt"
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	sub, err := nc.QueueSubscribeSync("tasks", "workers")
	if err != nil {
		fmt.Println(err)
		return
	}

	var msg *nats.Msg
	for {
		msg, err = sub.NextMsg(time.Hour * 10000)
		if err != nil {
			break
		}
		fmt.Println(string(msg.Data), msg.Reply, msg.Sub, msg.Subject, err)
	}
	nc.Flush()
	sub.Unsubscribe()
}
