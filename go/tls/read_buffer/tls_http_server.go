package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp4", ":1234")
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
				// TODO, cert and key
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
				resp := &http.Response{Header: http.Header{}, Body: http.NoBody}
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
