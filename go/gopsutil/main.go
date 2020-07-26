package main

import (
    "fmt"
	"os"
	"time"

    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

func cpuTest2() {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.CPUPercent())
}

func cpuTest() {
	rList, _ := cpu.Percent(time.Second, true)
	for _, r := range rList {
		fmt.Println(r)
    }
}

func memTest() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}

func main() {
	// cpuTest()
	// memTest()
	for {
		cpuTest2()
	}
}