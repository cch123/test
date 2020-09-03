package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestSymbol(t *testing.T) {
	cmd := exec.Command("nm", os.Args[0])
	contentBytes, err := cmd.Output()
	if err != nil {
		println(err)
		t.Fail()
	}

	fmt.Println(string(contentBytes))
	fmt.Println(os.Args[0])
}
