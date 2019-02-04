convert {123,23,1} to \uhex\uhex\uhex
-------

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// MyListener :
// java 用继承并 override 掉原来的 baseListener
// go 用内嵌原来的 baseListener，然后再实现这个 struct 想要实现的方法
type MyListener struct {
	BaseArrayInitListener
}

// EnterInit : 相当于 override 掉的原来的方法
func (ml *MyListener) EnterInit(c *InitContext) {
	fmt.Printf(`"`)
}

// ExitInit : 相当于 override 掉的原来的方法
func (ml *MyListener) ExitInit(c *InitContext) {
	fmt.Printf(`"`)
}

// EnterValue : 相当于 override 掉的原来的方法
func (ml *MyListener) EnterValue(c *ValueContext) {
	val, _ := strconv.ParseInt(c.INT().GetText(), 10, 64)
	res := fmt.Sprintf("%0[1]*[2]x", 4, val)
	fmt.Printf("\\u%v", res)
}

func main() {
	var input = antlr.NewInputStream("{132,2,5}")

	var lexer = NewArrayInitLexer(input)

	var tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewArrayInitParser(tokens)

	var tree = parser.Init()

	var listener = new(MyListener)
	var walker = antlr.NewParseTreeWalker()
	walker.Walk(listener, tree)

}
```