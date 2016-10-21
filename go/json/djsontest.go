package main

//这个库似乎只支持未知json格式的解析
//不支持decode到struct
import (
	"fmt"

	"github.com/a8m/djson"
)

type Person struct {
	name string
	age  int
}

func main() {
	//var p Person
	x, err := djson.Decode([]byte(`123`))
	fmt.Println(x, err)
	fmt.Printf("%q", x)
}
