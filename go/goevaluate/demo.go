package main

import (
	"fmt"
	"reflect"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, err := govaluate.NewEvaluableExpression("1+2+3 == 4 ? 1 : 0")
	v, err := expr.Eval(nil)
	fmt.Println(reflect.TypeOf(v), v, err)

	expr, err = govaluate.NewEvaluableExpression(`x == "abc" || x == "def" ? 1 : 0`)
	v, err = expr.Evaluate(map[string]interface{}{
		"x": "abc",
	})
	fmt.Println(v, err)

	expr, err = govaluate.NewEvaluableExpression("x & 1024 > 0")
	v, err = expr.Evaluate(map[string]interface{}{
		"x": 124,
	})
	fmt.Println(reflect.TypeOf(v), v, err)
}
