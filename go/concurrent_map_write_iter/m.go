package main

import (
	"sync"
)

var x map[int]string = make(map[int]string)

func f(s string, wg *sync.WaitGroup) {
	x[0] = s
	wg.Done()
}

func g(s string, wg *sync.WaitGroup) {
	x[1] = s
	wg.Done()
}

func main() {
	for {
		var wg sync.WaitGroup
		wg.Add(2)
		go f("Hello", &wg)
		go g("Playground", &wg)
		wg.Wait()
	}
}
