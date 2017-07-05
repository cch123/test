package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.OpenFile("./tmp/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	ch := make(chan struct{}, 100)
	go func() {
		for {
			ch <- struct{}{}
		}
	}()

	for i := 0; i < 5000; i++ {
		go func() {
			for {
				select {
				case _ = <-ch:
				}
			}
		}()
	}

	go func() {
		time.Sleep(time.Second * 10)
		pprof.StopCPUProfile()
		f.Close()
	}()
	http.ListenAndServe(":8080", nil)

}
