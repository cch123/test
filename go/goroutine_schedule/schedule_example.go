package main

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for v := range data {
			// 非内联函数，会被编译器插入调度代码
			println(v)
		}

		println("receive over")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)
	//runtime.Gosched()

	println("send over")
	<-exit
}
