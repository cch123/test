package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "1927-12-31 23:54:07")
	t2, _ := time.Parse("2006-01-02 15:04:05", "1927-12-31 23:54:08")
	println(t.Unix())
	println(t2.Unix())
	fmt.Println(t, t2)
}
