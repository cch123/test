package main

//go tool objdump可以看生成的汇编代码，
//然后在汇编内容里搜索gotoolobjdump.go
//可以找到对应的行数和汇编代码的关系

func main() {
	a := []int{1, 2, 3}
	b := len(a)
	println(b)
}
