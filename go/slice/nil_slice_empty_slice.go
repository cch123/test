package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    var a = []int{}
    var b []int = nil

    pa := (*reflect.SliceHeader)(unsafe.Pointer(&a))
    pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))

    fmt.Printf("a: %#v\n", *pa)
    fmt.Printf("b: %#v\n", *pb)
}
