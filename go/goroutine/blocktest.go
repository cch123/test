package main

func test(ch chan int) {
	ch <- 1
}

func main() {
	//第二个参数如果不提供的话，那么默认是阻塞
	//阻塞的channel是没法在一个goroutine中又发送又接收的
	//否则会死锁
	//例子
	/**
	  下面这种会死锁
	  var ch = make(chan int)
	  ch <- 1
	  x := <-ch
	  下面这种会死锁
	  var ch = make(chan int)
	  test(ch)
	  x:=<-ch
	  下面这种不会死锁
	  var ch = make(chan int)
	  go test(ch)
	  x:=<-ch
	  下面这种不会死锁
	  var ch = make(chan int, 1)
	  ch <-1
	  x:=<-ch
	  但是如果往buffered channel里塞过多的数据的话
	  var ch = make(chan int, 1)
	  ch <-1
	  ch <-1 //在这里阻塞了
	  x:=<-ch
	*/
	var ch = make(chan int, 1)
	test(ch) //error!
	//go test(ch)
	x := <-ch
	println(x)
}
