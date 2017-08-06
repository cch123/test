package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 这种场景可以用 gas/gometalinter 检查出来
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	fmt.Println(req)
}
