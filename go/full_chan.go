package main

func main() {
	var a = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case a <- i:
		default:
			println("in default", i)
		}
	}

}
