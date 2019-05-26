package main

import "sync/atomic"

var a int64

func atomicStore(i int) {
	atomic.StoreInt64(&a, int64(i))
}