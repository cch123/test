package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"a": []string{"1"},
		"b": []string{"2"},
		"c": []string{"3"},
		"d": []string{"4"},
	}

	// 会把所有的 k, v 输出出来
	for k, v := range m {
		fmt.Println(k, v)
		m = nil
		m = make(map[string][]string)
		m["b"] = []string{"x"}

		// m = nil
	}

}
