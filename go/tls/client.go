package main

import (
	"crypto/tls"
	"fmt"
)

func main() {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("/hello\n"))
	if err != nil {
		fmt.Println(err, n)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err, n)
		return
	}

	println(string(buf[:n]))
}
