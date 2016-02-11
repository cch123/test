package main

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func main() {
	var ch = make(chan int)
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	go sum(arr[0:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)
	var x, y = <-ch, <-ch
	println(x, y)
}
