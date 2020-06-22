package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/cch123/elasticsql"
)

var (
	tickInterval = time.Second * 2
	//memoryThreshold = 1000000000
	memoryThreshold = 10000
	rateThreshold   = 2.0 // 200%
)

type profileStats struct {
	totalAllocSize, totalAllocObject int
	totalFreedSize, totalFreedObject int
	totalInuseSize, totalInuseObject int
}

func init() {
	go func() {
		for {
			var a = make([]int, 100000)
			_ = a
			elasticsql.Convert("fuck")
		}
	}()
}

func main() {
	ticker := time.NewTicker(tickInterval)
	prev := profileStats{}
	for range ticker.C {
		curr := calculateTheCurrentMemprofile()
		fmt.Printf("%#v\n", curr)
		if toDump(prev, curr) {
			// dump
			writeHeapToFile()
		}
		prev = curr
	}
}

func writeHeapToFile() {
	suffix := time.Now().Format("2006-01-02.15:04:05")
	fileName := fmt.Sprintf("heap.%v", suffix)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[autodump] create file failed", err)
		return
	}

	err = pprof.Lookup("heap").WriteTo(f, 1)
	if err != nil {
		fmt.Println("[autodump] write profile to file failed", err)
	}
}

func toDump(prev, current profileStats) bool {
	// diff the two
	// if the thresholds are matched
	// dump
	if current.totalInuseSize > memoryThreshold {
		return true
	}

	return false
}

// returns alloc, freed, inuse
func calculateTheCurrentMemprofile() profileStats {
	var (
		totalAllocSize, totalAllocObject = 0, 0
		totalFreedSize, totalFreedObject = 0, 0
		totalInuseSize, totalInuseObject = 0, 0
	)

	profiles := make([]runtime.MemProfileRecord, 2)
	n, ok := runtime.MemProfile(profiles, false)
	if ok {
		profiles = profiles[0:n]
	} else {
		return profileStats{}
	}

	for _, profile := range profiles {
		totalAllocSize += int(profile.AllocBytes)
		totalFreedSize += int(profile.FreeBytes)
		totalInuseSize += int(profile.InUseBytes())

		totalAllocObject += int(profile.AllocObjects)
		totalFreedObject += int(profile.FreeObjects)
		totalInuseObject += int(profile.InUseObjects())
	}

	return profileStats{
		totalAllocSize:   totalAllocSize,
		totalAllocObject: totalAllocObject,
		totalFreedSize:   totalFreedSize,
		totalFreedObject: totalFreedObject,
		totalInuseSize:   totalInuseSize,
		totalInuseObject: totalInuseObject,
	}
}
