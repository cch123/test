package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func setLogLevel(level string) {
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
}

func setLogPath(path string) {
	rl := rotatelogs.New(path+".%Y%m%d",
		rotatelogs.WithMaxAge(720*time.Hour),
	)
	logrus.SetOutput(rl)
}

func main() {
	setLogPath("./logname")
	//setLogPath("/Users/caochunhui/test/go/log/name")
	logrus.Errorln("yes")
	logrus.Infoln("log rotation test")
}
