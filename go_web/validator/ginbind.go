// 注意gin里带的这个binding
// 是自定义了tag为binding的go-playground/validator组件
// 原生的插件默认是用validate这个tag的
// required 因为golang语义缺失的问题，只能检验某个字段是否是0值
// 这个一定要注意
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"
)

type person struct {
	Age  int    `json:"age" binding:"gte=0,lte=130"`
	Name int    `json:"name" binding:"lt=10"`
	AAA  string `json:"aaa" binding:"lt=10"`
	T    int    `json:"t" binding:"ne=1,ne=0"`
}

func contentType(r *http.Request) string {
	return filterFlags(requestHeader(r, "Content-Type"))
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func requestHeader(r *http.Request, key string) string {
	if values, ok := r.Header[key]; ok {
		return values[0]
	}
	return ""
}

func bind(r *http.Request, dst interface{}) error {
	contentType := contentType(r)
	if len(contentType) == 0 {
		// do some thing?
	}

	b := binding.Default(r.Method, contentType)
	return b.Bind(r, dst)
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	p := person{}
	err := bind(r, &p)
	fmt.Printf("%#v", p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
