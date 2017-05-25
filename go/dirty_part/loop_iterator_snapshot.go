package main

import (
	"fmt"
)

type s struct {
	S string
}

func main() {
	var s1 *s
	var s2 *s

	var a []s = []s{
		s{
			S: "aaa",
		},
		s{
			S: "bbb",
		},
	}

	for k, v := range a {
		fmt.Printf("%p\t%s\n", &v, (&v).S)
		if k == 0 {
			s1 = &v
		} else if k == 1 {
			s2 = &v
		}
	}

	fmt.Printf("s1: %p\tv: %s\n", s1, s1.S)
	fmt.Printf("s2: %p\tv: %s\n", s2, s2.S)
}
