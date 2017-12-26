package main

import (
    "reflect"
    "runtime"
)

func main() {
    fv := reflect.ValueOf(func(int) struct{} { return struct{}{} })
    args := []reflect.Value{reflect.ValueOf(0)}

    go func() {
        runtime.GC()
    }()

    for {
        fv.Call(args)
    }

}
