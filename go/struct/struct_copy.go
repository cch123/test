package main

import (
	"errors"
	"fmt"
	"reflect"
)

func copy(dst interface{}, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		err := errors.New("dst isn't a pointer to struct")
		return err
	}
	dstElem := dstVal.Elem()
	if dstElem.Kind() != reflect.Struct {
		err := errors.New("pointer doesn't point to struct")
		return err
	}

	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		return errors.New("src must be struct")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcType.Field(i)
		innerSrcVal := srcVal.FieldByName(srcField.Name)

		if innerDstVal := dstElem.FieldByName(srcField.Name); innerDstVal.IsValid() && innerDstVal.CanSet() {
			innerDstVal.Set(innerSrcVal)
		}
	}

	return nil
}

type A struct {
	Age    int
	Name   string
	Height int
}

type B struct {
	Age  int
	Name string
}

func main() {
	var a = A{}
	var b = B{2343, "abc"}

	copy(&a, b)
	fmt.Println(a)
}
