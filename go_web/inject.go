package main

import "github.com/codegangsta/inject"
import "fmt"

/*
type injector struct {
    values map[reflect.Type]reflect.Value
    parent Injector
}

键值对是：
key=>变量类型
value=>变量值
所以参数列表里有类型相同的内容的话，会发生覆盖
不太科学吧
*/

//这两个类型都是为了防止发生覆盖声明的自己的interface类型
//为了在interfaceOf函数调用时能得到不同的字符串作map的key
type SpecialString interface{}
type MyInt interface{}

func say(name string, gender SpecialString, age MyInt, length int) {
	fmt.Printf("My name is %s, gender is %s, age is %d, length is %d!\n", name, gender, age, length)
}

func main() {
	inj := inject.New()
	inj.Map("Xargin")
	inj.MapTo("男", (*SpecialString)(nil))
	//inj.MapTo("男", (*SpecialString)(nil))
	//inj.MapTo(10, (*MyInt)(nil))
	inj.MapTo(10, (*MyInt)(nil))
	inj.Map(1000)
	inj.Invoke(say)
}
