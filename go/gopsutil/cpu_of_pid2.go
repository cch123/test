package main

import (
	"fmt"
	"runtime"
	"os"
//	"strconv"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
go func () {for {}}()
	//pid, _ := strconv.ParseInt(os.Args[1], 10, 64)
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		percent, err:= p.CPUPercent()
		fmt.Println(percent, err)
		fmt.Println(int(percent)/runtime.GOMAXPROCS(-1))
		time.Sleep(time.Second)
	}
}
