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
	"time"

	"encoding/json"

	"github.com/cch123/binding"
)

type person struct {
	X int       `json:"x" form:"x" binding:"required,gt=-1,lt=10"`
	Y *int      `json:"y" form:"y" binding:"required,gt=-1,lt=10"`
	Z time.Time `json:"z" form:"z" binding:"required"`
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
	fmt.Println(contentType)

	b := binding.Default(r.Method, contentType)
	return b.Bind(r, dst)
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	p := person{}
	err := bind(r, &p)
	fmt.Printf("%#v\n", p)
	fmt.Println("err is ", err)
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
