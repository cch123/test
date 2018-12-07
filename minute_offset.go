package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	todayZeroTime := time.Date(nowTime.Year(),
		nowTime.Month(), nowTime.Day(),
		0, 0, 0, 0, time.Local)

	offset := time.Now().Sub(todayZeroTime) / time.Duration(time.Second*60)
	fmt.Println(int(offset))
	fmt.Println(nowTime)
	fmt.Println(todayZeroTime)
}
