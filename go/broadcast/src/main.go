package main

import (
	"broadcast"
	"fmt"
	"time"
)

var b = broadcast.NewBroadcaster()

func listen(r broadcast.Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		go listen(r)
		fmt.Println(v)
	}
}

func main() {
	r := b.Listen()
	go listen(r)
	for i := 0; i < 10; i++ {
		b.Write(i)
	}
	b.Write(nil)

	time.Sleep(3 * 1e9)
}
