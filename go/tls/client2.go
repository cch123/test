package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"go.uber.org/ratelimit"
)

func main() {
	_, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	qps, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	bucket := ratelimit.New(int(qps))

	for i := 0; ; i++ {
		bucket.Take()
		go func() {
			conn := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
					IdleConnTimeout:     0,
					MaxIdleConns:        1,
					MaxIdleConnsPerHost: 1,
				},
			}
			defer conn.CloseIdleConnections()
			if resp, e := conn.Get("https://localhost:8443"); e != nil {
				fmt.Println(e)
			} else {
				defer resp.Body.Close()
				ioutil.ReadAll(resp.Body)
			}
		}()
	}

}
