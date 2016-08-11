package main

import (
	"fmt"
	"log"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://localhost:2379"},
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)
	var ctx = context.Background()
	resp, err := kapi.Set(ctx, "test", "bar", nil)
	if err != nil {
		if err == context.Canceled {
			// ctx is canceled by another routine
		} else if err == context.DeadlineExceeded {
			// ctx is attached with a deadline and it exceeded
		} else if cerr, ok := err.(*client.ClusterError); ok {
			// process (cerr.Errors)
			fmt.Println(cerr)
		} else {
			// bad cluster endpoints, which are not etcd servers
		}
	}
	fmt.Printf("%q\n", resp)
}
