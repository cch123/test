package main

import (
	"sort"
)

func test1() {
	twoSum1([]int{5, 4, 3, 2, 1}, 9)
}

func test2() {
	twoSum2([]int{5, 4, 3, 2, 1}, 9)
}

func twoSum1(nums []int, target int) []int {
	back := make([]int, len(nums))
	copy(back, nums)
	sort.Ints(nums)
	length := len(nums)
	var i, idx int
	for i = 0; i < length-1; i++ {
		t := target - nums[i]
		idx = sort.SearchInts(nums[i+1:], t)
		if i+idx+1 < length && nums[i+idx+1] == t {
			break
		}
	}
	a1, a2 := nums[i], nums[i+idx+1]
	var ansa, ansb int
	for j := 0; j < length; j++ {
		if back[j] == a1 || back[j] == a2 {
			if ansa == 0 {
				ansa = j
			} else {
				ansb = j
				break
			}
		}
	}
	if ansa > ansb {
		ansa, ansb = ansb, ansa
	}
	return []int{ansa, ansb}
}

func twoSum2(nums []int, target int) []int {
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
