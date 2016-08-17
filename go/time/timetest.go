package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Unix(1469179423000/1000, 0)
	fmt.Println(time)
}
