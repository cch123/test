package main

import (
	"fmt"
	"log"
	"runtime"
)

func Trace() {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s    line=%d\n", file, line)
	}
}

func main() {
	Trace()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("here here")
}
