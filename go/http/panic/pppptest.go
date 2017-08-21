package main

import (
	"fmt"
	"time"
)

func firePanic() {
	defer rescue()
	panic(1)
}

func rescue() {
	if err := recover(); err != nil {
		fmt.Println("oh fuck, panicked", err)
	}
}

func say(s string) {
}

func main() {

	// At the end of main, call rescue function
	defer rescue()

	// An amazing app starts here
	say("hello world")
	go firePanic()
	say("never reached, if panics are managed")
	time.Sleep(time.Second)

}
