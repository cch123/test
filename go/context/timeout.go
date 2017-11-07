package main

import (
	"context"
	"fmt"
	"time"
)

func test(ctx context.Context) {
	// must support ctx
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
		}
		fmt.Println("do some thing")
		time.Sleep(time.Millisecond)
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*4)
	test(ctx)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("yesyes")
		cancel()
	case <-ctx.Done():
		fmt.Println("nono")
	}
	fmt.Println(ctx)
}
