package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = map[string]interface{}{
		"a": "abc",
		"b": map[string]string{
			"d": "d",
			"c": "c",
		},
	}

	var b = map[string]interface{}{
		"b": map[string]string{
			"c": "c",
			"d": "d",
		},
		"a": "abc",
	}

	var x = reflect.DeepEqual(a, b)
	fmt.Println(x)

}
