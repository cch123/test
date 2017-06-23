package main

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	println(len(ch))
	println(cap(ch))
	<-ch
	println(len(ch))
	println(cap(ch))
}
