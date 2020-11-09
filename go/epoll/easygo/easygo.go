package main

import (
	"bufio"
	"fmt"
	"github.com/mailru/easygo/netpoll"
	"io"
	"net"
	"net/http"
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

		f, err := c.(*net.TCPConn).File()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(f.Fd())
		desc := netpoll.NewDesc(f.Fd(), netpoll.EventRead)
		err = poller.Start(desc,
			func(ev netpoll.Event) {
				fmt.Println("event come, event go", ev)
				fmt.Println(ev.String())

				req, err := http.ReadRequest(bufio.NewReader(f))
				if err == io.EOF {
					fmt.Println("read EOF")
					defer f.Close()
					poller.Stop(desc)
					return
				}

				if err != nil {
					fmt.Println(err, "read request failed")
					return
				}

				fmt.Println(req)
				resp := http.Response{
					StatusCode : http.StatusOK,
					ProtoMajor: 1,
					ProtoMinor: 1,
				}
				err = resp.Write(f)
				if err != nil {
					fmt.Println(err, "write failed")
					return
				}
			},
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
