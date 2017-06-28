package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zcalusic/sysinfo"
)

func main() {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
