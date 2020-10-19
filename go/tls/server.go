package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
)

func main() {

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world"))
		if err != nil {
			fmt.Println(err, n)
			return
		}
	}
}
