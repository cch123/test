package broadcast

import "fmt"

type broadcastChan chan broadcast

type broadcast struct {
	c broadcastChan
	v interface{}
}

type Broadcaster struct {
	// private fields:
	Listenc chan chan broadcastChan
	Sendc   chan<- interface{}
}

type Receiver struct {
	// private fields:
	C broadcastChan
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenc := make(chan chan broadcastChan)
	sendc := make(chan interface{})
	go func() {
		currc := make(broadcastChan, 1)
		for {
			select {
			case v := <-sendc:
				// write的时候会到这个case
				// 只listen的话其实只走下面那个分支
				fmt.Println("secret is not here")
				if v == nil {
					currc <- broadcast{}
					return
				}
				// 这里生成新的currc
				// 要往这里面塞进对应的数据
				// 这里的b的channel指向
				c := make(broadcastChan, 1)
				b := broadcast{c: c, v: v}
				currc <- b
				currc = c
			case r := <-listenc:
				fmt.Println("secret is here")
				r <- currc
			}
		}
	}()
	return Broadcaster{
		Listenc: listenc,
		Sendc:   sendc,
	}
}

// start listening to the broadcasts.
func (b Broadcaster) Listen() Receiver {
	c := make(chan broadcastChan, 0)
	// 重点是这里，从channel里取了出来，实际上又塞进去一个东西
	b.Listenc <- c
	// 这里不会block，为啥？
	r := Receiver{<-c}

	// 再加一个就block了
	// r = Receiver{<-c}
	// 所以实际上是在b.Listenc <- c之后，
	// case r := <-listenc:
	// r <- currc
	// 从listenc里取出了这个c，然后再把当前/下一个currc塞了进去

	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", r)
	// 这里不会block，为啥？

	return r
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{}) { b.Sendc <- v }

// read a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
	var b broadcast = <-r.C //这里拿到的是一个broadcast对象
	var v interface{} = b.v //broadcast对象里的v interface{}
	r.C <- b                //因为值取出来了，r.C里现在没东西
	fmt.Println(r.C)
	r.C = b.c
	fmt.Println(r.C)
	fmt.Println(b.c)
	return v
}
