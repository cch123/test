package main

import (
	"fmt"
	"regexp"
	"strings"
)

// golang里的split和java里的不一样
// java里的String str; str.split("\\.")需要转义，默认是正则
// 但在go里，如果需要正则形式split的话，需要用到regexp包
func main() {
	fmt.Println(strings.Split("a.b.c", `.`))
	reg := regexp.MustCompile("\\.") // 和java的split等价
	fmt.Println(reg.Split("a.b.c", -1))
}
