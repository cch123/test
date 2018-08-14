package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
)

func watchAndUpdate() {
}

func set() error {
	return nil
}

func get() (string, error) {
	return "", nil
}

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	w := kapi.Watcher("/name", nil)
	go func() {
		for {
			resp, err := w.Next(context.Background())
			fmt.Println(resp, err)
			fmt.Println("new values is ", resp.Node.Value)
		}
	}()

	log.Print("Setting /name to alex")
	resp, err := kapi.Set(context.Background(), "/name", "alex", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
	}

	log.Print("Getting /name key value")
	resp, err = kapi.Get(context.Background(), "/name", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get is done. Metadata is %q\n", resp)
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
	time.Sleep(time.Minute)
}
