package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//	"github.com/cheekybits/genny/generic"

var maxRetryTimes = 2

type taskHandler func(context.Context) (int, error)

type task struct {
	handler taskHandler
	timeout time.Duration
}

type runner struct {
	//taskChannel chan task //generic.Type
	complete   chan error
	tasks      []task
	resChannel chan int
}

func (r *runner) run(ctx context.Context, timeout time.Duration) {
	var wg = &sync.WaitGroup{}
	wg.Add(len(r.tasks))
	for _, t := range r.tasks {
		t := t // avoid range variable capture
		go func() {
			for i := 0; i < maxRetryTimes; i++ {
				res, err := t.handler(ctx)
				if err != nil {
					continue
				}
				// err == nil
			}
		}()
	}
	wg.Wait()
}

func main() {
	var r = runner{}
	r.complete = make(chan error, 1)
	r.resChannel = make(chan int, 10)
	r.tasks = []task{
		task{
			handler: func(ctx context.Context) (int, error) {
				time.Sleep(time.Second)
				return 1, nil
			},
			timeout: time.Millisecond * 100,
		},
	}
	r.run(context.TODO(), time.Millisecond*10)
	select {
	case <-r.complete:
		fmt.Println("all complete!")
	}
}
