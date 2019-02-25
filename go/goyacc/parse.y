%{
package main

import (
    "fmt"
    "text/scanner"
    "os"
    "strings"
)

const NEQStr = "!="
type Expression interface{}

type BinOpExpr struct {
    left     Expression
    operator string
    right    Expression
}
%}

%union{
    token string
    expr  Expression
    neq string
}

%type<expr> program
%type<expr> expr
%token<token> NUMBER
%token<token> NEQ

%left '+' NEQ

%%

program
    : expr
    {
        $$ = $1 // 会把 $1 当成返回值赋值给 return val
        yylex.(*Lexer).ast= $$
    }

expr
    : NUMBER
    {
        $$ = $1
    }
    | NEQ
    {
        $$ = NEQStr
    }
    | expr NEQ expr
    {
        $$ = BinOpExpr{left: $1, operator: "!=", right: $3}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{left: $1, operator: "+", right: $3}
    }
%%

type Lexer struct {
    scanner.Scanner
    ast Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
    token := int(l.Scan())
    if token == scanner.Int {
        token = NUMBER
    }
    if l.TokenText() == "!" {
        token = NEQ
        l.Scan()
        lval.token = "!="
    } else {
        lval.token = l.TokenText()
    }
    return token
}

func (l *Lexer) Error(e string) {
    panic(e)
}

func main() {
    l := new(Lexer)
    l.Init(strings.NewReader(os.Args[1]))
    yyParse(l)
    fmt.Printf("%#v\n", l.ast)
}
