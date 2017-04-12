package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {

		if _, hasFound := m[target-v]; hasFound {
			left := k
			right := m[target-v]
			if left > right {
				left, right = right, left
			}
			return []int{left, right}

		}

		m[v] = k
		//append(m[v], k)
	}

	return []int{0, 0}
}
