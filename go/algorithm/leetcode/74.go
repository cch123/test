package main

import "fmt"

func main() {
	var matrix = [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 13, 14}, []int{23, 30, 44, 55}}
	searchMatrix(matrix, 3)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	var i, j = 0, len(matrix[0]) - 1
	for {
		// i j is valid
		if i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) {
		} else {
			fmt.Println("here", i, j)
			return false
		}

		if matrix[i][j] > target {
			j--
			continue
		}

		if matrix[i][j] < target {
			i++
			continue
		}

		if matrix[i][j] == target {
			return true
		}
	}
}
