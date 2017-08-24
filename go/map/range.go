package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
	}

	for k, v := range m {
		fmt.Printf("m ptr %p\n", m)
		fmt.Println(k, v)
		m = nil
		fmt.Printf("m ptr %p\n", m)
	}

	time.Sleep(time.Hour)

}
