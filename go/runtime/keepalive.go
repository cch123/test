package main

import (
	"runtime"
	"time"
)

type Person struct {
	age int
}

func getPerson() *Person {
	var p = &Person{11}
	runtime.SetFinalizer(p, func(x *Person) {
		println("gc happen on x", x)
	})
	return p
}

/*
keep alive 主要用在下面这种场景:
f, _ := os.Open()
fd := f.Fd()

syscall.Read(fd)

// 确保 fd 在 Read 返回前一直都是有效的，因为 f 已经相当于没有引用了
runtime.Keepalie(fd)
*/

func main() {
	var p = getPerson()
	_ = p
	runtime.GC()
	runtime.GC()
	// try to comment out the keepalive and try again
	runtime.KeepAlive(p)
	time.Sleep(time.Minute)
}
