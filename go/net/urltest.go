package main

//import "fmt"
import "net/url"
import "encoding/json"

//import "strings"

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	println(s)
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	x, err := json.Marshal(u)
	println(string(x))

	s = "http://www.baidu.com/fucker"
	println(s)
	u, err = url.Parse(s)
	if err != nil {
		panic(err)
	}
	x, err = json.Marshal(u)
	println(string(x))

}
