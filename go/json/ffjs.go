package cch

//ffjson这个是用已有的struct自动生成Marshal和Unmarshal的代码的库
//必须得把代码放到GOPATH下才能正常运行。。
//执行ffjson xxxxx.go，然后可以生成编解码代码

type Person struct {
	name string
	age  int
}
