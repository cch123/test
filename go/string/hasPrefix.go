package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "g_qc_source_0"
	fmt.Println(strings.HasPrefix(a, "g_qc_source"))
	fmt.Println(strings.HasPrefix(a, "g_qc_sourcw"))
	fmt.Println(fmt.Sprintf("%v_%v", 1, "1"))

}
