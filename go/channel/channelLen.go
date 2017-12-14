package main

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 1
	for i := 0; i < 100; i++ {

		go func() {
			<-ch
		}()
	}
	ch <- 1
	ch <- 1
	println(len(ch))
	println(cap(ch))
	<-ch
	println(len(ch))
	println(cap(ch))
}
