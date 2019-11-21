package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		lf = flag.String("l", "", "flag local")
		rf = flag.String("r", "", "flag remote")
	)

	flag.Parse()
	if len(*rf) == 0 || len(*lf) == 0 {
		flag.Usage()
		return
	}

	ra, err := net.ResolveTCPAddr("tcp4", *rf)
	if err != nil {
		fmt.Println(err)
		return
	}

	la, err := net.ResolveTCPAddr("tcp4", *lf)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", la)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		remoteConn, err := net.DialTCP("tcp4", nil, ra)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("connect from origin : %v, connecting to remote : %v\n", conn.RemoteAddr().String(), *rf)

		go func() {
			defer func() {
				conn.Close()
				remoteConn.Close()
			}()

			var (
				inboundChan  = make(chan struct{}, 1)
				outBoundCHan = make(chan struct{}, 1)
			)

			go func() {
				pipe(conn, remoteConn)
				inboundChan <- struct{}{}
			}()

			go func() {
				pipe(remoteConn, conn)
				outBoundCHan <- struct{}{}
			}()

			// 任意一个 g 退出，都应该直接关闭两条连接
			select {
			case <-inboundChan:
				println("quit from inbound")
				return
			case <-outBoundCHan:
				println("quit from outbound")
				return
			}
		}()

	}

}

func pipe(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
	}
}
