package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"bou.ke/monkey"
)

func main() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})

	fmt.Println("what the hell?") // what the *bleep*?
	monkey.UnpatchAll()
	t()
	monkey.Patch(http.Post, func(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
		fmt.Println("oh yes this is the patched http post")
		return nil, nil
	})
	http.Post("abcdefg", "adfsdf", nil)
}

func t() {
	fmt.Println("hello world")
}
