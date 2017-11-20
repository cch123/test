package main

import (
	"context"
	"fmt"
	"time"
)

func test(ctx context.Context) {
	fmt.Println("slow start")
	time.Sleep(time.Second * 2)
	fmt.Println("slow done")
}

func main() {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	go test(ctx)
	select {
	case <-time.After(time.Second):
		fmt.Println("yesyes")
		cancel()
	case <-ctx.Done():
		fmt.Println("nono")
	}
	fmt.Println(ctx)
	time.Sleep(time.Second * 2)
}
