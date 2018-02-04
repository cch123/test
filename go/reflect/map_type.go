package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a map[string]int
	typ := reflect.TypeOf(a)
	kind := reflect.TypeOf(a).Kind()
	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	elemType := reflect.TypeOf(a).Elem()
	// error !!! not addressable
	// reflect.ValueOf(a).Elem().Set(reflect.MakeMap(reflect.TypeOf(a)))
	// can use pointer to set elem
	reflect.ValueOf(&a).Elem().Set(reflect.MakeMap(reflect.TypeOf(a)))
	reflect.ValueOf(a).SetMapIndex(reflect.ValueOf("abc"), reflect.ValueOf(1))
	keyType := reflect.TypeOf(a).Key()
	fmt.Println(typ)
	fmt.Println(kind)
	fmt.Println(elemType)
	fmt.Println(keyType)
	fmt.Println(a)
}
