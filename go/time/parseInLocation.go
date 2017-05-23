package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	rateDate, _ := time.ParseInLocation("2006-01-02", a.Format("2006-01-02"), time.Local)
	fmt.Println(rateDate)
	//fmt.Printf("%#v", time.Local)
	var x = 10
	var b = time.Minute * time.Duration(100)
	fmt.Println(rateDate.Add(b))

}
