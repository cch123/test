package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/levigross/grequests"
)

//httpbin.org是个不错的网站，可以拿来做测试用
//不过国内怎么这么慢啊。。。。

func main() {
	//resp, err := grequests.Get("http://httpbin.org/get", nil)
	//fmt.Println(resp)
	//fmt.Println(err)
	var a = map[string]int{
		"aaa": 1,
		"bbb": 1,
	}
	by, _ := json.Marshal(a)

	var b = bytes.NewReader(by)
	///var a = strings.NewReader("aaa")

	ro := &grequests.RequestOptions{
		RequestBody: b,
		Headers:     map[string]string{"Content-Type": "application/json"},
	}
	resp, err := grequests.Post("http://httpbin.org/post", ro)
	fmt.Println(resp)
	fmt.Println(err)
}
