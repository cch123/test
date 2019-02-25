%{
package main

import (
    "fmt"
    "text/scanner"
    "os"
    "strings"
)

type Expression interface{}
type Token struct {
    token   int
    literal string
}

type NumExpr struct {
    literal string
}
type BinOpExpr struct {
    left     Expression
    operator rune
    right    Expression
}
%}

%union{
    token Token
    expr  Expression
}

%type<expr> program
%type<expr> expr
%token<token> NUMBER

%left '+'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

expr
    : NUMBER
    {
        $$ = NumExpr{literal: $1.literal}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{left: $1, operator: '+', right: $3}
    }

%%

type Lexer struct {
    scanner.Scanner
    result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
    token := int(l.Scan())
    if token == scanner.Int {
        token = NUMBER
    }
    lval.token = Token{token: token, literal: l.TokenText()}
    return token
}

func (l *Lexer) Error(e string) {
    panic(e)
}

func main() {
    l := new(Lexer)
    l.Init(strings.NewReader(os.Args[1]))
    yyParse(l)
    fmt.Printf("%#v\n", l.result)
}
