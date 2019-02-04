```
grun ArrayInit init -gui
```

```
grun ArrayInit init -tokens
```

```
{1,3,{1,5,32},32,{233,3}}
```


```go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	var input = antlr.NewInputStream("{1,2,{44,43},5}")

	var lexer = NewArrayInitLexer(input)

	var tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewArrayInitParser(tokens)

	var tree = parser.Init()

    // 这个 ToStringTree 在内部根本没用到 []string{} 这个变量
    // 看来 golang 版还不是很完善
	fmt.Println(tree.ToStringTree([]string{}, parser))

}

```