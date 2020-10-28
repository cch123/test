package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	pid, _ := strconv.ParseInt(os.Args[1], 10, 64)
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println(p.CPUPercent())
		time.Sleep(time.Second)
	}
}
