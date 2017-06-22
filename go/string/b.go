package main

func main() {
	var a = map[string]int{"abc": 1}
	var b = []byte{'a', 'b', 'c'}
	// runtime.slicebytetostring
	var c = string(b)
	_ = a[c]

	// 这里会被优化掉
	// 不会出现runtime.slicebytetostring
	_ = a[string(b)]
}
