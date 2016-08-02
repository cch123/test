package main

//第i行的位置j
var queen_idx = [8]int{}

//统计出了几种解法
var counter = 0

func main() {
	eight_queen(0)
	println(counter)
}

func check_valid(row_idx, col_val int) bool {
	flag := true
	for i := 0; i < row_idx; i++ {
		if queen_idx[i] == col_val {
			flag = false
			break
		}
		if my_abs(queen_idx[i]-col_val) == my_abs(row_idx-i) {
			flag = false
			break
		}
	}
	return flag
}

func eight_queen(row int) {
	for i := 0; i < 8; i++ {
		if check_valid(row, i) {
			queen_idx[row] = i
			if row == 7 {
				for k, v := range queen_idx {
					println(k, v)
				}
				println()
				counter++
				return
			}
			eight_queen(row + 1)
		}
	}
}

func my_abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
