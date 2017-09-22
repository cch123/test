package main

import (
	"fmt"
	"time"
)

func main() {
	//	a := time.Now()
	rateDate, _ := time.ParseInLocation("2006-01-02", "2015-01-02", time.Local)

	fmt.Println(rateDate, rateDate.Format("20060102 15:04:05"))
	r, _ := time.Parse("20060102", "20150102")
	fmt.Println(r, r.Format("20060102 15:04:05"))
	//fmt.Printf("%#v", time.Local)
	//	var x = 10
	var b = time.Minute * time.Duration(100)
	fmt.Println(rateDate.Add(b))

}
