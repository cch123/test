package main

import (
	"context"
	"fmt"
	"reflect"

	"bou.ke/monkey"
)

type lightBalancer struct{}

func (lb *lightBalancer) no(ctx context.Context, key string) (string, error) {
	return "", nil
}

func main() {
	lb := &lightBalancer{}

	monkey.PatchInstanceMethod(reflect.TypeOf(lb), "no", func(_ *lightBalancer, ctx context.Context, key string) (string, error) {
		return "", fmt.Errorf("no dialing allowed")
	})

	lb.no(context.TODO(), "")
}
