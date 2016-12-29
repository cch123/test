package main

import "fmt"
import "sort"

type A struct {
	a int
}

// implement these 3 methods to implement the sort.Interface
func (s ASlice) Len() int {
	return len(s)
}

func (s ASlice) Less(i, j int) bool {
	return s[i].a < s[j].a
}

// slice是引用传递的，注意
func (s ASlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//func (s ASlice)

type ASlice []A

func main() {
	// how to get the pointer of this new type ?
	var a = ASlice{{1}, {0}}
	fmt.Println(a)
	fmt.Printf("%p\n", a)
	sort.Sort(a)
	fmt.Printf("%p\n", a)
	fmt.Println(a)
}
