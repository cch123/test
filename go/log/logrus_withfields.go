package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	fmt.Println(log.GetLevel())
	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.SetLevel(log.DebugLevel)
	fmt.Println(log.GetLevel())
}
