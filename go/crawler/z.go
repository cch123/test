package main

import (
	"fmt"
	"net/url"
)

func main() {

	u, err := url.Parse("http://www.abc.com/afdsaf")
	fmt.Println(u.Hostname(), err)
}
