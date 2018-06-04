package main

//go:noinline
func sliceParam(sl []int64) {
}

//go:noinline
func stringParam(str string) {
}

//go:noinline
func arrayParam(sl [3]int) {
}

//go:noinline
func chanParam(ch chan int) {
}

//go:noinline
func mapParam(m map[int]int) {
}

func main() {
	ch := make(chan int, 10)
	sliceParam([]int64{1, 2, 3,6,7, 10,11})
	stringParam("abcdefg")
	arrayParam([3]int{4, 6,5})
	chanParam(ch)
	mapParam(map[int]int{})
}
