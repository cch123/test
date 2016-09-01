package main

import (
	"fmt"
	"time"
)

func main() {
	time0 := time.Unix(1469179423000/1000, 0)
	time1 := time.Unix(1479189423000/1000, 0)
	time2 := time.Unix(1479189423000/1000, 0)
	diff := time1.Sub(time0)
	diff2 := time2.Sub(time1)

	fmt.Println(diff)
	if diff > time.Hour*7*24 {
		fmt.Println("time diff bigger than 7 days")
	}

	fmt.Println(diff2)
	if diff2 >= time.Hour*7*24 {
		fmt.Println("time diff less bigger than 7 days")
	}

	//get current time
	fmt.Println(time.Now())
	fmt.Println(time0)
	fmt.Println(time1)
	fmt.Println(time2)
}
