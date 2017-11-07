package main

import (
	"fmt"
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
	c := &http.Client{
		Timeout: time.Millisecond * 2000,
	}
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	getEnd := time.Since(getStart)
	fmt.Println("get cost", getEnd.String())
}
