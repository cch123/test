package main

import "fmt"

func main() {
	var a = []int{1, 2, 4, 5, 6, 0}
	merge(a, 5, []int{3}, 1)
	fmt.Println(a)
}

//merge to nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	var j = m - 1
	var k = n - 1
	var i = m + n - 1
	for i >= 0 {
		if j < 0 || k < 0 {
			break
		}

		if nums1[j] < nums2[k] {
			nums1[i] = nums2[k]
			k--
			i--
		} else {
			nums1[i] = nums1[j]
			j--
			i--
		}
	}

	for j >= 0 {
		nums1[i] = nums1[j]
		j--
		i--
	}

	for k >= 0 {
		nums1[i] = nums2[k]
		k--
		i--
	}
}
