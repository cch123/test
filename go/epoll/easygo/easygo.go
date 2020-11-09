package main

import (
	"fmt"
	"net"

	"github.com/mailru/easygo/netpoll"
)

func main() {
	poller, err := netpoll.New(&netpoll.Config{
		OnWaitError: func(err error) {
			fmt.Println("wait failed", err)
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.Listen("tcp4", ":10101")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		tcpConn := c.(*net.TCPConn)
		f, err := tcpConn.File()
		if err != nil {
			fmt.Println(err)
			continue
		}

		var buf = make([]byte, 1024)
		var off = 0
		err = poller.Start(netpoll.NewDesc(f.Fd(), netpoll.EventRead),
			func(ev netpoll.Event) {

				fmt.Println("event come, event go", ev)
				fmt.Println(ev.String())

				n, err := tcpConn.Read(buf[off:])
				fmt.Println("tcp read", n, err, string(buf))

				if off+n == cap(buf) {
					bufNew := make([]byte, cap(buf)*2)
					copy(bufNew, buf)
				}

				off = n
			},
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
