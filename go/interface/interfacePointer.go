package main

import "bytes"
import "fmt"
import "io"

func main() {
	var a *bytes.Buffer
	fmt.Println(a == nil)
	f(a)
	g(a)
}

func f(out io.Writer) {
	if out != nil {
		fmt.Println("not nil in f")
		//out.Write([]byte("fuck"))
		return
	}
	fmt.Println("nil in f")
}

func g(out *bytes.Buffer) {
	if out != nil {
		fmt.Println("not nil in g")
		//out.Write([]byte("fuck"))
		return
	}
	fmt.Println("nil in g")
}
