package main

import (
	"fmt"
)

const MAX = 10

var sortArray = []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}

func main() {

	fmt.Println("before sort：")
	show()

	quickSort(sortArray, 0, MAX-1)

	fmt.Println("after sort:")
	show()

}

func quickSort(sortArray []int, left, right int) {
	if left < right {
		pos := partition(sortArray, left, right)
		partition(sortArray, left, pos-1)
		partition(sortArray, pos+1, right)
	}
}

func partition(sortArray []int, left, right int) int {
	key := sortArray[right]
	i := left - 1

	for j := left; j < right; j++ {
		if sortArray[j] <= key {
			i++
			swap(i, j)
		}
	}

	swap(i+1, right)

	return i + 1
}

// Swap the position of a and b
func swap(a, b int) {
	sortArray[a], sortArray[b] = sortArray[b], sortArray[a]
}

// foreach
func show() {
	for _, value := range sortArray {
		fmt.Printf("%d\t", value)
	}
}
