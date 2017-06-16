package main

import "strings"

func A() []string {
	var a = "aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\neee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t"
	var res = make([]string, 0, 36)
	startOffset := 0
	endOffset := 0
	for i := 0; i < len(a); i++ {
		endOffset = i
		if a[i] == '\t' || a[i] == '\n' {
			strSlice := a[startOffset:endOffset]
			startOffset = endOffset + 1
			//println(string(strSlice), strSlice, len(strSlice))
			res = append(res, strSlice)
		}
	}
	return res
}

func B() []string {
	var a = "aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\neee\tfff\tggg\thhh\tiii\t" +
		"aaa\tbbb\tccc\tddd\teee\tfff\tggg\thhh\tiii\t"
	a = strings.Replace(a, "\n", "\t", -1)
	b := strings.Split(a, "\t")
	return b
}

func main() {
	A()
}
