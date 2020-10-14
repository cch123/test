package main

import (
	"io"
	"sync"
)

var chunkPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1<<14)
	},
}

type bufChain struct {
	head *node
	tail *node
}

type node struct {
	chunk   [1 << 14]byte // 16 KB
	next    *node
	prevous *node
}

type mySpecialReader struct {
	chain     bufChain
	bufIndex  int
	bufOffset int
}

var randomDataBuffers = [][]byte{
	make([]byte, 1<<2),
	make([]byte, 1<<4),
	make([]byte, 1<<8),
	make([]byte, 1<<16),
	make([]byte, 1<<17),
	make([]byte, 1<<20),
	make([]byte, 1<<22),
}

// 这个函数的作用就是从 r 里读内容，如果把 m 的 chain 里的 buffer 读满了，就扩容
func (m *mySpecialReader) Read(r io.Reader) (int, error) {
	return 0, nil
}

func main() {
}
