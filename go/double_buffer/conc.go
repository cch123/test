package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type config struct {
	fk2Opts    map[string]string
	topic2Opts map[string]string
	ev2Opts    map[string]string
	l          sync.RWMutex
}
type doublebuffer struct {
	c1         config
	c2         config
	currentIdx int64
}

var option doublebuffer

func readConfig() (bool, bool, bool) {
	idx := atomic.LoadInt64(&option.currentIdx)
	if idx == 0 {
		option.c1.l.RLock()
		defer option.c1.l.RUnlock()
		_, ok1 := option.c1.fk2Opts["1"]
		_, ok2 := option.c1.topic2Opts["x"]
		_, ok3 := option.c1.ev2Opts["1"]
		return ok1, ok2, ok3
	}

	option.c2.l.RLock()
	defer option.c2.l.RUnlock()
	_, ok1 := option.c2.fk2Opts["1"]
	_, ok2 := option.c2.topic2Opts["1"]
	_, ok3 := option.c2.ev2Opts["x"]
	return ok1, ok2, ok3
}

func writeToConfig() {
	fk2Opts := map[string]string{}
	topic2Opts := map[string]string{}
	ev2Opts := map[string]string{}
	for i := 0; i < 10000; i++ {
		fk2Opts[fmt.Sprint(i)] = fmt.Sprint(i)
		topic2Opts[fmt.Sprint(i)] = fmt.Sprint(i)
		ev2Opts[fmt.Sprint(i)] = fmt.Sprint(i)
	}
	idx := atomic.LoadInt64(&option.currentIdx)
	newVal := 1 - idx
	if newVal == 0 {
		option.c1.l.Lock()
		defer option.c1.l.Unlock()
		option.c1.fk2Opts = fk2Opts
		option.c1.topic2Opts = topic2Opts
		option.c1.ev2Opts = ev2Opts
	} else {
		option.c2.l.Lock()
		defer option.c2.l.Unlock()
		option.c2.fk2Opts = fk2Opts
		option.c2.topic2Opts = topic2Opts
		option.c2.ev2Opts = ev2Opts
	}

	atomic.CompareAndSwapInt64(&option.currentIdx, idx, 1-idx)
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			for {
				writeToConfig()
			}
		}()

		go func() {
			for {
				readConfig()
			}
		}()
	}
	time.Sleep(time.Hour)
}
