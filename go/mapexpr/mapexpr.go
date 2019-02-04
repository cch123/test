package mapexpr

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Eval : 判断 map 是否符合 bool 表达式
//	expr = `a > 1 && b < 0`
func Eval(m map[string]string, expr string) (bool, error) {
	//fset := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, err
	}

	// ast.Print(fset, exprAst)
	// TODO，深度优先搜索，判断 bool 正确与否
	return judge(exprAst, m), nil
}

func isLeaf(bop ast.Node) bool {
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	_, okL := expr.X.(*ast.Ident)
	_, okR := expr.Y.(*ast.BasicLit)
	if okL && okR {
		return true
	}
	return false
}

// dfs
func judge(bop ast.Node, m map[string]string) bool {
	// binary expression 才能直接返回 true false
	// 否则返回的都是递归的结果
	if isLeaf(bop) {
		// do the leaf logic
		expr := bop.(*ast.BinaryExpr)
		x := expr.X.(*ast.Ident)
		y := expr.Y.(*ast.BasicLit)
		return m[x.Name] == y.Value
	}

	// not leaf
	// 那么一定是 binary expression
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		println("this cannot be true")
		return false
	}

	switch expr.Op {
	case token.LAND:
		//println("op is &&")
		//println("left is", judge(expr.X, m))
		//println("right is", judge(expr.Y, m))
		return judge(expr.X, m) && judge(expr.Y, m)
	case token.LOR:
		//println("op is ||")
		//println("left is", judge(expr.X, m))
		//println("right is", judge(expr.Y, m))
		return judge(expr.X, m) || judge(expr.Y, m)
	}

	println("unsupported operator")
	return false
}
