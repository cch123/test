package main

import "os/exec"

var uuid []byte

func test() {
	uuid, _ = exec.Command("uuidgen").Output()
}
