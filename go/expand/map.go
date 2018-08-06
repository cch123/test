// map 即使在数组内部扩容了，也不会生成新的局部变量
package main

import "fmt"

func testMap(m map[int]int) {
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
}

func main() {
	var m = map[int]int{}
	testMap(m)
	fmt.Println(m)
}
