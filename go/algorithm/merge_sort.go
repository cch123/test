package main

func mergeSort(input []int) {
	// from interval 1 to interval len(input)/2
	// merge the array
	var interval = 1
	for interval < len(input) {
		// merge and merge
		var startIndex = 0
		for startIndex < len(input) {
			var leftSlice []int
			var rightSlice []int
			if startIndex+interval <= len(input) {
				leftSlice = input[startIndex : startIndex+interval]
				if startIndex+interval*2 <= len(input) {
					rightSlice = input[startIndex+interval : startIndex+interval*2]
				} else {
					rightSlice = input[startIndex+interval : len(input)]
				}

			} else {
				// 说明right不存在，不做任何处理
				// left的上界取len
				leftSlice = input[startIndex:len(input)]
			}

			// merge
			var merged = merge(leftSlice, rightSlice)

			for i := startIndex; i < startIndex+len(merged); i++ {
				input[i] = merged[i-startIndex]
			}
			startIndex += interval * 2
			//fmt.Println("left is ", leftSlice, "right is ", rightSlice)
		}

		interval *= 2
	}
	//fmt.Println(input)
}

func main() {
	mergeSort([]int{5, 4, 3, 2, 1, 10, 5, 3, 5, 6, 7, 8, 2, 8, 2, 22, 34})
}

func merge(left []int, right []int) []int {
	if len(right) == 0 {
		return left
	}

	var totalLength = len(left) + len(right)
	var merged = make([]int, totalLength)

	// ensure the left is the small array
	if len(left) > len(right) {
		left, right = right, left
	}

	var i, j, k = 0, 0, 0

	for {
		// one is done
		if i >= len(left) || j >= len(right) {
			break
		}

		if left[i] <= right[j] {
			merged[k] = left[i]
			//这一堆++，还是C或者C艹方便啊
			k++
			i++
		} else {
			merged[k] = right[j]
			k++
			j++
		}
	}

	// fetch the rest of left
	for i < len(left) {
		merged[k] = left[i]
		k++
		i++
	}

	// fetch the rest of right
	for j < len(right) {
		merged[k] = right[j]
		k++
		j++
	}
	return merged
}
