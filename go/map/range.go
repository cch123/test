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

	// 会把所有的 k, v 输出出来
	for k, v := range m {
		fmt.Println(k, v)
		m = nil
	}

	time.Sleep(time.Hour)

}
