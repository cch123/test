package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/panicwrap"
)

func main() {
	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		// Something went wrong setting up the panic wrapper. Unlikely,
		// but possible.
		panic(err)
	}

	// If exitStatus >= 0, then we're the parent process and the panicwrap
	// re-executed ourselves and completed. Just exit with the proper status.
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	// Otherwise, exitStatus < 0 means we're the child. Continue executing as
	// normal...

	// Let's say we panic
	//panic("oh shucks")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func panicHandler(output string) {
	// output contains the full output (including stack traces) of the
	// panic. Put it in a file or something.
	fmt.Printf("The child panicked:\n\n%s\n", output)
	panic(1)
	go func() {
		fmt.Println("enter goroutine")
		go func() {

			panic(1)
		}()
	}()
}
