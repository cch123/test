package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		hasModel = false
		path     = "/Users/didi/test"
	)

	/*
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// trigger
			goto TRIGGER
		}
	*/

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".model") {
			hasModel = true
		}
		return nil
	})
	println(err, hasModel)

	if !hasModel {
		goto TRIGGER
	} else {
		return
	}
TRIGGER:
	println("fuck")
}
