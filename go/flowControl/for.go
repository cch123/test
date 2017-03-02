package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
	/*
		it is the same with len = len(v)
		for i:=0;i<len;i++ {
		}
	*/
}
