package main

type Printer interface {
	Print()
}

type person struct{}

func (p person) Print() {}

type gogo struct{}

func (g gogo) Print() {}

func main() {

}
