package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {

	go func() {
		var sigs = make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		for s := range sigs {
			logrus.Infof("Got signal: %s, service stoped.", s)
			os.Exit(1)
		}
	}()

	for {
		time.Sleep(time.Second)
	}

}
