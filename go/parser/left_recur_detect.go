package main

import (
	"fmt"
	"os"
)

func main() {
	var ruleMap = map[string][]string{
		"expr":      []string{"and_expr", "or_expr", "comp_expr"},
		"comp_expr": []string{"field", "op", "value"},
		"and_expr":  []string{},
	}
	for k, rules := range ruleMap {
		recordMap := map[string]struct{}{
			k: struct{}{},
		}
		// TODO
		var nextLines = rules
		for len(nextLines) > 0 {
			var toBeAppend []string
			for _, l := range nextLines {
				if _, ok := recordMap[l]; ok {
					fmt.Println("left recursion exist!", l, recordMap)
					os.Exit(1)
				}
				recordMap[l] = struct{}{}
				toBeAppend = append(toBeAppend, rules...)
			}
		}
	}
	fmt.Println(ruleMap)
}
