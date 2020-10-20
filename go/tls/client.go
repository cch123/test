package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/juju/ratelimit"
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

	bucket := ratelimit.NewBucket(time.Second/time.Duration(qps), qps)

	connList := []*http.Client{}
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

	for i := 0; ; i++ {
		bucket.Wait(1)
		conn := connList[i%len(connList)]
		go func() {
			if resp, e := conn.Get("https://localhost:4443"); e != nil {
				fmt.Println(e)
			} else {
				defer resp.Body.Close()
				io.Copy(ioutil.Discard, resp.Body)
			}
		}()
	}

}
