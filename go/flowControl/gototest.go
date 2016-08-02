package main

func test(i int) {
	if i < 0 {
		goto ERROR_HANDLE
	}
	return
ERROR_HANDLE:
	println("this is error handle")
	return
}

func main() {
	println("test 2")
	test(2)
	println("test -1")
	test(-1)
}
