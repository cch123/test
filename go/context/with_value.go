package main

import (
	"context"
	"fmt"
)

type orderID int

func main() {
	var x = context.TODO()
	x = context.WithValue(x, orderID(1), "1234")
	x = context.WithValue(x, orderID(2), "2345")
	y := context.WithValue(x, orderID(3), "4567")
	x = context.WithValue(x, orderID(3), "3456")
	fmt.Println(x.Value(orderID(3)))
	fmt.Println(y.Value(orderID(3)))
}
