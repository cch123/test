package main

import "io"
import "net"
import "time"

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("get msg from server:\n", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/cch.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	go reader(c)
	for {
		_, err := c.Write([]byte("hello world"))
		if err != nil {
			println("write error", err)
			break
		}
		time.Sleep(1e9)
	}
}
