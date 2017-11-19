package main

import (
	"fmt"
	"net"

	"github.com/bsm/redeo"
	"github.com/bsm/redeo/resp"
)

func main() {
	// Init server and define handlers
	srv := redeo.NewServer(nil)

	srv.HandleFunc("lpush", func(w resp.ResponseWriter, cmd *resp.Command) {
		if cmd.ArgN() < 2 {
			w.AppendError("lacking param")
			return
		}

		msgQueue := cmd.Arg(0)
		msg := cmd.Arg(1)
		// use this msg to do some msg queue stuff
		fmt.Println("queue: ", msgQueue, ", msg: ", msg)

		// ... succeed, then give response
		w.AppendOK()
	})

	srv.HandleFunc("rpop", func(w resp.ResponseWriter, cmd *resp.Command) {
		// msg := fetch one msg from your message queue
		fakeMsg := "yes, this msg is fetched from msg queue"
		w.AppendInlineString(fakeMsg)
	})

	// Open a new listener
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	srv.Serve(lis)
}
