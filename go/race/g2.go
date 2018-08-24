package main

import "runtime"

var a int

func calc() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				runtime.RaceDisable()
				a++
				runtime.RaceEnable()
			}
		}()

	}
}

func main() {
	calc()
}
