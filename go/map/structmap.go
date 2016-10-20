package main

import "fmt"

//map的key可以是string以外的值
//所以也可以用来做指针和对应状态的记录
//不过如果是和json相关的工作时，需要考虑到rfc的json定义里object的key只能是string
//int类型做key的map会出错

type Node struct {
	length int
	next   *Node
}

func main() {
	var a = make(map[*Node]bool)
	n := Node{}
	a[&n] = true
	fmt.Println(a)
}
