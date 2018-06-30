package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://baidu.com")
	defer func() {
		fmt.Printf("defer %#v\n", resp.Body)
		resp.Body.Close()
	}()
	fmt.Printf("before %#v\n", resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	fmt.Printf("after %#v\n", resp.Body)
}
