package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 2}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	//var left, right = 0, 0
	var left, right = 0, len(height) - 1
	var min = height[left]
	if height[right] < height[left] {
		min = height[right]
	}
	var maxArea = min * (right - left)
	// 贪心规则
	// 一小一大，要移动小的一方，否则可用反证得到该移动方法不是正确的
	for left < right {

		min = height[left]
		if height[right] < height[left] {
			min = height[right]
		}

		var area = min * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}
