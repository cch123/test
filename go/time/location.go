//export TZ=Asia/Tokyo go run location.go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location:", t.Location(), ":Time:", t)
	utc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("Location:", utc, ":Time:", t.In(utc))

}
