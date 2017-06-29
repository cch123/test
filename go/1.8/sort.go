package main

import (
	"fmt"
	"sort"
)

type Person struct {
	age  int
	name string
}

func main() {
	var a = []Person{Person{1, "alex"}, Person{0, "nono"}}
	// 1.8 之后支持这样 sort
	// 非常方便
	sort.Slice(a, func(i, j int) bool { return a[i].age < a[j].age })
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool { return a[i].name < a[j].name })
	fmt.Println(a)
}
