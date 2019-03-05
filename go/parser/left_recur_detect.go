package main

import "fmt"

func main() {
	var ruleMap = map[string][]string{
		"expr":      []string{"and_expr", "or_expr", "comp_expr"},
		"comp_expr": []string{"field", "op", "value"},
		"and_expr":  []string{},
	}
	fmt.Println(ruleMap)
}
