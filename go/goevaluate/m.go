package main

import (
	"fmt"
	"reflect"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, err := govaluate.NewEvaluableExpression("1+2+3")
	v, err := expr.Eval(nil)
	fmt.Println(reflect.TypeOf(v), v, err)
}
