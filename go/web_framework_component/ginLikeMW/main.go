package main

import "fmt"

func main() {
	c := NewContext()

	var sign string
	c.use(func(c *MyContext) {
		sign += "a"
		c.next()
		sign += "b"
	})

	c.use(func(c *MyContext) {
		sign += "c"
		c.next()
		sign += "d"
	})

	c.use(func(c *MyContext) {
		sign += " fuck "
	})

	c.startHandle()
	fmt.Println(sign)
}
