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

type Person struct {
	Age  int    `json:"age" binding:"gte=0,lte=130"`
	Name int    `json:"name" binding:"required,lt=10"`
	AAA  string `json:"aaa" binding:"required"`
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	p := Person{}
	fmt.Println(r.Header["Content-Type"])
	contentType := r.Header["Content-Type"]
	if len(contentType) > 0 {
		b := binding.Default("POST", contentType[0])
		err := b.Bind(r, &p)
		fmt.Println(p, err)
	} else {
		fmt.Println("no content type")
	}
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
