package main

import "github.com/coreos/pkg/capnslog"
import "os"

var plog = capnslog.NewPackageLogger("soundwave", "main")

func main() {
	capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stdout, false))
	capnslog.SetGlobalLogLevel(capnslog.TRACE)
	plog.Debug("aaa")
	plog.Trace("aaa")
	plog.Info("aaa")
	plog.Notice("aaa")
	plog.Warning("aaa")
	plog.Error("aaa")
	plog.Fatal("aaa")
	plog.Panic("aaa")
}
