package main

import (
	"fmt"
	"time"

	"context"
	"os"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

//https://github.com/coreos/etcd/blob/ace3a217b00c537c059df066f7f262669f900faa/integration/v3_lock_test.go

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		os.Exit(1)
	}

	defer cli.Close()
	sess, err := concurrency.NewSession(cli)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mutex := concurrency.NewMutex(sess, "xxxxb")
	err = mutex.Lock(ctx)

	println("lock acquired")
	fmt.Println(err)
	defer mutex.Unlock(ctx)
	for {
		println("in process")
		time.Sleep(time.Second)
	}

}
