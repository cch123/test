package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"go.uber.org/ratelimit"
)

func main() {
	connNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	qps, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := //os.Args[3]
		"http://100.88.107.242:8080"

	bucket := ratelimit.New(int(qps))

	var l sync.Mutex
	connList := make([]*http.Client, connNum)
	/*
		for i := 0; i < int(connNum); i++ {
			connList = append(connList, &http.Client{
				Transport: &http.Transport{
					TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
					IdleConnTimeout:     0,
					MaxIdleConns:        1,
					MaxIdleConnsPerHost: 1,
				},
			})
		}
	*/

	for i := 0; ; i++ {
		bucket.Take()
		i := i
		go func() {
			l.Lock()
			if connList[i%len(connList)] == nil {
				connList[i%len(connList)] = &http.Client{
					Transport: &http.Transport{
						IdleConnTimeout:     0,
						MaxIdleConns:        1,
						MaxIdleConnsPerHost: 1,
					},
				}
			}
			conn := connList[i%len(connList)]
			l.Unlock()
			if resp, e := conn.Get(url); e != nil {
				fmt.Println(e)
			} else {
				defer resp.Body.Close()
				ioutil.ReadAll(resp.Body)
			}
		}()
	}

}
