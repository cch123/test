package main

import "sync"

var nm = make(map[int]int)
var nmLock = sync.RWMutex{}
var sm = sync.Map{}

func normalMapInsert(k, v int) {
	nmLock.Lock()
	nm[k] = v
	nmLock.Unlock()
}

func normalMapGet(k int) int {
	nmLock.RLock()
	defer nmLock.RUnlock()
	return nm[k]
}

func syncMapInsert(k, v int) {
	sm.Store(k, v)
}

func syncMapGet(k int) int {
	res, ok := sm.Load(k)
	if ok {
		return res.(int)
	}
	return 0
}