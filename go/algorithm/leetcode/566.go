package main

func main() {
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || r*c != len(nums)*len(nums[0]) {
		return nums
	}

	//valid
	var i, j, ii, jj int
	var res [][]int
	for ii < r {
		var tmpRow = make([]int, c)
		for jj = 0; jj < c; jj++ {
			tmpRow[jj] = nums[i][j]
			if j == len(nums[0])-1 {
				j = 0
				i++
			} else {
				j++
			}
		}
		jj = 0
		ii++
		res = append(res, tmpRow)
	}
	return res
}
