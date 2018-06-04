package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ret []string
	args := []interface{}{string("001"), string("003"), string("004")}
	cp(args, &ret)
	fmt.Println(ret)
}

func cp(rep interface{}, ret interface{}) {
	fmt.Println(reflect.TypeOf(rep))
	fmt.Println(reflect.ValueOf(rep).Kind())
	retVal := reflect.ValueOf(ret).Elem()
	srcVal := reflect.ValueOf(rep)
	if reflect.ValueOf(rep).Kind() == reflect.Slice {
		retSl := reflect.MakeSlice(retVal.Type(), srcVal.Len(), srcVal.Cap())
		for i := 0; i < retSl.Len(); i++ {
			srcIValInter := srcVal.Index(i)
			// interface{} 不能直接赋给 string
			// 必须要用 elem 取
			retSl.Index(i).Set(srcIValInter.Elem())
		}

		fmt.Println("assignable", srcVal.Type().AssignableTo(retVal.Type()))
		fmt.Println("assignable", srcVal.Type(), retVal.Type())
		reflect.ValueOf(ret).Elem().Set(retSl)
	}
}
