package main

func main() {
	var arr = []int{1, 2, 3}
	for k, v := range arr {
		println(k, v)
	}
	arr = append(arr, 4, 5, 6)
}
