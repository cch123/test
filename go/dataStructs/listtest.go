package main

import "container/list"
import "fmt"

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		fmt.Println(elem.Value)
	}
}
