package main

import "sync/atomic"

var id int64

func t() {
	atomic.AddInt64(&id, 1)
}

func task(i int) {
	println(i)
	for {
		//如果在for里不调用不可内联的函数，那就会4个goroutine阻塞在这里
		//调用的话，会被编译器在函数调用前插入调度指令
		t()
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		go task(i)
		atomic.AddInt64(&id, 1)
	}
	ch := make(chan struct{})
	<-ch
}
