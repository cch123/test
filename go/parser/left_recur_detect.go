package main

import "fmt"

func main() {
	var ruleMap = map[string][]string{
		"expr":      []string{"and_expr", "or_expr", "comp_expr"},
		"comp_expr": []string{"field", "op", "value"},
		"and_expr":  []string{},
	}
	for k, rules := range ruleMap {
		recordMap := map[string]struct{}{}
		insertRec := func(k string) {
			recordMap[k] = struct{}{}
		}
		// TODO
	}
	fmt.Println(ruleMap)
}

func traverseNextLine(ruleMap map[string][]string, insert func(string)) bool {
}
