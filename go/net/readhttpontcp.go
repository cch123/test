package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

var respContent = `HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: max-age=86400
Connection: Keep-Alive
Content-Length: 0
Content-Type: text/html
Date: Sat, 11 Aug 2018 10:17:47 GMT
ETag: "51-47cf7e6ee8400"
Expires: Sun, 12 Aug 2018 10:17:47 GMT
Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
Server: Apache

`

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":13000")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go func() {
			r, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				fmt.Println("fuck", err)
				return
			}
			fmt.Printf("%#v\n", r)
			conn.Write([]byte(respContent))
			fmt.Println(r.ContentLength)
		}()
	}
}
