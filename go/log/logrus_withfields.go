package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
