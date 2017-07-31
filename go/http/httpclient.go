package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://baidu.com")
	fmt.Println(resp)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	fmt.Println(err)
}
