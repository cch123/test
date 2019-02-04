```
 antlr4 -no-listener -visitor -Dlanguage=Go LabeledExpr.g4
```

```
package main

import (
	"fmt"
	"strconv"

	"github.com/wxio/antlr4/runtime/Go/antlr"
)

// EvalVisitor 继承 base visitor，实现自定义的 visitor 方法
type EvalVisitor struct {
	BaseLabeledExprVisitor
}

var memory = make(map[string]interface{})

func (ev *EvalVisitor) VisitAssign(ctx *AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := ev.Visit(ctx.Expr())
	memory[id] = value
	return value
}

func (ev *EvalVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	value := ev.Visit(ctx.Expr())
	fmt.Println(value)
	return 0
}

func (ev *EvalVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	left := ev.Visit(ctx.Expr(0))
	right := ev.Visit(ctx.Expr(1))
	if ctx.op.GetTokenType() == LabeledExprParserMUL {
		// 没有泛型，这里写起来好麻烦
		return left.(int) * right.(int)
	}
	return left.(int) / right.(int)
}

func (ev *EvalVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	left := ev.Visit(ctx.Expr(0))
	right := ev.Visit(ctx.Expr(1))
	if ctx.op.GetTokenType() == LabeledExprParserADD {
		// 没有泛型，这里写起来好麻烦
		return left.(int) + right.(int)
	}
	return left.(int) - right.(int)
}

func (ev *EvalVisitor) VisitId(ctx *IdContext) interface{} {
	id := ctx.ID().GetText()
	if res, ok := memory[id]; ok {
		return res
	}
	return 0
}

func (ev *EvalVisitor) VisitInt(ctx *IntContext) interface{} {
	value, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	return int(value)
}

func (ev *EvalVisitor) VisitParens(ctx *ParensContext) interface{} {
	return ev.Visit(ctx.Expr())
}

var str = `
193
a = 4
b = 5
a+b*2
(1+2)*3
`

func main() {
	var input = antlr.NewInputStream(str)

	var lexer = NewLabeledExprLexer(input)

	var tokens = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := NewLabeledExprParser(tokens)

	var tree = parser.Prog()

	var eval = new(EvalVisitor)
	eval.Visit(tree)

}

```
程序理论上应该是这样的

但是 antlr 目前的 runtime/Go 不支持 visitor pattern

https://github.com/antlr/antlr4/pull/1841

https://github.com/antlr/antlr4/pull/1807

目前蛋疼中