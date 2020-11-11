package main

import "github.com/tidwall/evio"

func main() {
	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		// out 其实就是写出的数据，这里 out = in，实现的其实是一个 echo server
		out = in
		return
	}
	if err := evio.Serve(events, "tcp://localhost:5000"); err != nil {
		panic(err)
	}
}
