package main

import "time"

func main() {
	var a []string
	for {
		var d = len(a)

		time.Sleep(time.Second)
		println(d)
	}
}
