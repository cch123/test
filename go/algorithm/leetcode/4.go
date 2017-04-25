package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 处理单数组为空的情况

	// 未处理单数组为空的情况
	evenFlag = false
	if (len(nums1)+len(nums2))%2 == 0 {
		evenFlag = true
	}

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numMerged := make([]int, len(nums1)+len(nums2))
	i, j, k := 0, 0, 0
	for {
		if i >= len(nums1) || j >= len(nums2) {
			break
		}

		if nums1[i] < nums2[j] {
			numMerged[k] = nums1[i]
			i++
		} else {
			numMerged[k] = nums2[j]
			j++
		}

		k++
	}

	for i < len(nums1) {
		numMerged[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) {
		numMerged[k] = nums2[j]
		j++
		k++
	}

	if len(numMerged)%2 == 0 {
		return float64(numMerged[(len(numMerged)-1)/2])/2 + float64(numMerged[len(numMerged)/2])/2
	}
	return float64(numMerged[len(numMerged)/2])
}
