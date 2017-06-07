package main

import (
	"reflect"
	"unsafe"
)

func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func TestC() {
	var a = []byte{'a', 'b'}
	var b = map[string]bool{}
	b[String(a)] = true
}

func TestA() {
	var a = []byte{'a', 'b'}
	var b = map[string]bool{}
	b[string(a)] = true
}

func TestB() {
	var a = []byte{'a', 'b'}
	var b = map[string]bool{}
	var c = string(a)
	b[c] = true
}
