package main

// MyContext is ss
type MyContext struct {
	handlers []Handler
	index    int
}

// NewContext is
func NewContext() *MyContext {
	return &MyContext{index: 0}
}

// Handler is
type Handler func(c *MyContext)

func (c *MyContext) next() {
	c.index++
	length := len(c.handlers)
	for ; c.index < length; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *MyContext) use(h Handler) {
	c.handlers = append(c.handlers, h)
}

func (c *MyContext) startHandle() {
	c.handlers[0](c)
}
