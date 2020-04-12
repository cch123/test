package main

type p struct{}

func (pp p) Get() {
}

func (pp p) Set() {
}

type X interface {
	Get()
}

type Y interface {
	Set()
}

func main() {
	var x interface{} = p{}
	switch x.(type) {
	case Y:
		println("yyy")
	case X:
		println("xxx")
	default:
		println("cannot be true")
	}
}
