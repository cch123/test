%{
package main

import (
    "fmt"
    "text/scanner"
    "os"
    "strings"
)

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
    bin_op string
}

%type<expr> program
%type<expr> expr
%type<bin_op> bin_op

%token<token> NUMBER

//%token<bin_op> NEQ GTE LTE

// 下面的 NEQ 和 lexer 中返回的 token 应该是对应的
// 有运算符优先级的定义的话
// 似乎也不需要在上面的 token 进行定义了
%left '+' '-' NEQ GTE LTE
%left '*' '/'

%%

program
    : expr
    {
        $$ = $1 // 会把 $1 当成返回值赋值给 return val
        yylex.(*Lexer).ast= $$
    }

bin_op
    : NEQ { $$ = "!="}
    | GTE { $$ = ">=" }
    | LTE { $$ = "<=" }
    | '+' { $$ = "+" }
    | '-' { $$ = "-" }
    | '*' { $$ = "*" }

expr
    : NUMBER
    {
        $$ = $1
    }
    | expr bin_op expr
    {
        $$ = BinOpExpr{left: $1, operator: $2, right: $3}
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
    // 这里需要额外处理一些多字符的 token 的情况
    // 比如 ! 开头的
    // 比如 a in [1,2,3,4 这种的]
    // 比如 a is null 这种的
    // >= <= 等等
    if l.TokenText() == "!" {
        token = NEQ
        l.Scan()
        lval.token = "!="
    } else if l.TokenText() ==">" {
        token = GTE
        l.Scan()
        lval.token = ">="
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
