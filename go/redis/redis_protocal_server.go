package main

import (
	"net"

	"github.com/bsm/redeo"
	"github.com/bsm/redeo/resp"
)

func main() {
	// Init server and define handlers
	srv := redeo.NewServer(nil)
	srv.HandleFunc("ping", func(w resp.ResponseWriter, _ *resp.Command) {
		w.AppendInlineString("PONG")
	})
	srv.HandleFunc("info", func(w resp.ResponseWriter, _ *resp.Command) {
		w.AppendBulkString(srv.Info().String())
	})
	srv.HandleFunc("set", func(w resp.ResponseWriter, _ *resp.Command) {
		w.AppendArrayLen(2)
		w.AppendInlineString("a")
		w.AppendInt(5)
	})

	// Open a new listener
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	// Start serving (blocking)
	srv.Serve(lis)
}
