package main

import (
	"fmt"
	"time"

	"github.com/ngaut/go-zookeeper/zk"
)

// NewKazoo creates a new connection instance
var conn *zk.Conn
var err error

func NewKazoo(servers []string) {

	conn, _, err = zk.Connect(servers, time.Second*5)
	if err != nil {
		println(err)
		return
	}
	fmt.Printf("%#v\n", conn)
}

func WatchInstances() <-chan zk.Event {
	fmt.Println("watch instances here")

	node := fmt.Sprintf("%s/consumers/%s/ids", "", "offline-test")
	if exists, _, err := conn.Exists(node); err != nil {
		fmt.Println(err)
	} else if !exists {
		fmt.Println("node not exist")
	}

	_, _, c, err := conn.ChildrenW(node)
	if err != nil {
		return nil
	}

	return c
}

func main() {
	var a = []string{"10.94.112.43:2181"}
	NewKazoo(a)
	c := WatchInstances()
	for {
		select {
		case <-c:
			fmt.Println("consumer change")
		default:
			time.Sleep(time.Second)
		}
	}
}
