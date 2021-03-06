package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe(":9999", nil)
}

func main() {
	l, err := net.Listen("tcp4", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			c = tls.Server(c, &tls.Config{
				Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true,
			})

			if err != nil {
				fmt.Println(err)
				return
			}

			r := bufio.NewReader(c)
			for {
				req, err := http.ReadRequest(r)
				if err != nil {
					fmt.Println(err)
					c.Close()
					return
				}

				_, err = ioutil.ReadAll(req.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				// write respose
				resp := &http.Response{ProtoMajor: 1, ProtoMinor: 1,StatusCode : http.StatusOK, Header: http.Header{}, Body: http.NoBody}
				err = resp.Write(c)
				if err != nil {
					fmt.Println(err)
					c.Close()
					return
				}
			}
		}()
	}
}
