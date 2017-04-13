package main

func main() {
	var a = make(chan int, 4)
	close(a)
	// 直接panic，思考一下why
	a <- 4
}
