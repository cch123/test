package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	mainStart := time.Now()
	//n := 10
	//n := 100
	n := 1000
	rsyncGet(n)
	mainEnd := time.Since(mainStart)
	fmt.Println("main cost", mainEnd.String())
}
func rsyncGet(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			get()
		}()
	}
	wg.Wait()
}

func get() {
	getStart := time.Now()
	resp, _ := http.Get("http://www.baidu.com")
	io.Copy(ioutil.Discard, resp.Body)

	defer resp.Body.Close()
	getEnd := time.Since(getStart)
	fmt.Println("get cost", getEnd.String())
}
