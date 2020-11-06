package main

import (
	"bytes"
	"fmt"
	"runtime/pprof"
	"strings"
	"unicode"
)

func main() {
	var buf bytes.Buffer
	pprof.Lookup("threadcreate").WriteTo(&buf, 1)

	l, _ := buf.ReadString('\n')
	l = strings.TrimSpace(l)
	var (
		i            = len(l) - 1
		threadNumStr string
	)

	for unicode.IsDigit(rune(l[i])) && i >= 0 {
		threadNumStr = string([]byte{l[i]}) + threadNumStr
		i--
	}
	fmt.Println(threadNumStr)
}
