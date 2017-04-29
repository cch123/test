package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			sum += nums[i]
		}
	}
	return sum
}
