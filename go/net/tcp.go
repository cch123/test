package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", ":3456")
	l, err := net.ListenTCP("tcp4", addr)
	defer l.Close()

	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			defer conn.Close()

			_, err := conn.Write([]byte("hello world\n"))
			if err != nil {
				fmt.Println(err)
				return
			}

			for {
				var buf = make([]byte, 1024)
				size, err := conn.Read(buf)
				if err != nil {
					conn.Close()
					break
				}
				buf = buf[:size]
				fmt.Printf("%s", string(buf))
				size, err = conn.Write(buf)
				if err != nil {
					conn.Close()
					break
				}
				fmt.Println("write", size, "bytes")
			}
		}()
	}
}
