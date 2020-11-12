package main

import (
	"fmt"
	"github.com/xtaci/gaio"
	"net"
)

func echoServer(w *gaio.Watcher) {
	for {
		results, err := w.WaitIO()
		if err != nil {
			panic(err)
		}

		for _, res := range results {
			switch res.Operation {
			case gaio.OpRead:
				if res.Error == nil {
					w.Write(nil, res.Conn, res.Buffer[:res.Size])
				} else {
					fmt.Println("err when read", res.Error)
				}
			case gaio.OpWrite:
				if res.Error == nil {
					w.Read(nil, res.Conn, res.Buffer[:cap(res.Buffer)])
				}
			}
		}
	}
}

func main() {
	w, err := gaio.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer w.Close()

	go echoServer(w)

	ln, err := net.Listen("tcp", ":1999")
	if err != nil {
		panic(err)
	}

	fmt.Println("listening on ", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("accepted", conn.RemoteAddr())

		// submit async io read io req
		err = w.Read(nil, conn, make([]byte, 128))
		if err != nil {
			panic(err)
		}
	}
}
