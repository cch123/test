package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	res, err := c.Create("/platform/order-system/create-order-service-http/10.1.13.3:1043", []byte("1"),
		zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println(res)
	time.Sleep(time.Second * 50)

	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
