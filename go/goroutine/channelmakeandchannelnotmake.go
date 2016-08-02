package main

func main() {
	var a chan int
	var b = make(chan int)
	//注意，只声明不初始化的话，a的值是空，没法直接使用
	//go里的三种特殊数据结构都是需要初始化的
	println(a)
	println(b)
	a = make(chan int)
	println(a)
}
