package main

import "fmt"

func add(a int, b int) int
func minus(a int, b int) int
func sum(s []int64) int64
func sum1(s []int64) (ret int64) {
	for i := 0; i < len(s); i++ {
		ret += s[i]
	}
	return ret
}

func main() {
	arr := []int64{1, 2, 3, 4, 10}
	sux := sum(arr)
	fmt.Println(sux)
	fmt.Println(add(90, 101))
	fmt.Println(minus(1021, 323))
	fmt.Println(sum1(arr))

}
