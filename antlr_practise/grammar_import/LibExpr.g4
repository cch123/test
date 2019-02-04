grammar LibExpr;
import CommonLexerRules; // import 已经存在的语法规则
/* 起始规则 */
prog : stat + ; // 程序由多个 statement 组成

stat : expr NEWLINE
     | ID '=' expr NEWLINE
     | NEWLINE
     ;

expr : expr ('*' | '/') expr
     | expr ('+' | '-') expr
     | INT
     | ID
     | '(' expr ')'
     ;
