package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	var conns = []net.Conn{}
	var ff = []*os.File{}
	for {
		conn, err := l.Accept()
		conns = append(conns, conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		f, err := conn.(*net.TCPConn).File()
		if err != nil {
			fmt.Println(err)
			return
		}
		f1 := os.NewFile(f.Fd(), "")
		ff = append(ff, f1)
		fmt.Println("second copy of file", f1)
		runtime.GC()
	}
}
