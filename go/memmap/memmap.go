package main

import (
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/bradleyjkemp/memviz"
	"github.com/sitano/gsysint"
)

func main() {
	var gp unsafe.Pointer

	atomic.StorePointer(&gp, gsysint.GetG())

	gg := (*gsysint.G)(gp)
	memviz.Map(os.Stdout, gg)
}
