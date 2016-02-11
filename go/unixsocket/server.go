package main

import "net"

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println(string(data))
		_, err = c.Write(data)
		if err != nil {
			panic("Write: " + err.Error())
		}
	}
}

func main() {
	l, err := net.Listen("unix", "/tmp/cch.sock")
	if err != nil {
		println("listen error", err.Error())
		return
	}
	for {
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			return
		}
		go echoServer(fd)
	}
}
