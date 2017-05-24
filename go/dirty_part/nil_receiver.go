package main

import "log"
import "reflect"

type A struct {
	value string
}

func (a *A) Test() string {
	log.Print("call Test")
	log.Print(reflect.TypeOf(a))
	log.Print(reflect.ValueOf(a))
	return a.value
}

func getA() *A {
	return nil
}

func main() {
	a := getA()
	a.Test()
}
