package main

import (
	"fmt"
	"sync"
)

func dostuff(wg *sync.WaitGroup, i int) {
	fmt.Printf("goroutine id %d\n", i)
	fmt.Printf("goroutine id %d\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	workers := 10

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go dostuff(&wg, i)
	}
	wg.Wait()
}
