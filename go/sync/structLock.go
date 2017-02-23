package main

import "sync"

type Person struct {
	age  int
	lock sync.Mutex
}

// 考虑一下这里为什么不能用p Person?
func (p *Person) changeAge(newAge int) {
	p.lock.Lock()
	p.age = newAge
	p.lock.Unlock()
}

func main() {
	p := new(Person)
	p.changeAge(5)
}
