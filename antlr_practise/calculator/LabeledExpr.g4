grammar LabeledExpr;

/* 起始规则，语法分析的起点 */
prog : stat +;

// 注意，这些 expr 的标签会被用来生成代码的！
stat : expr NEWLINE # printExpr
     | ID '=' expr NEWLINE # assign
     | NEWLINE # blank
     ;

expr : expr op=('*'|'/') expr # MulDiv
     | expr op=('+'|'-') expr # AddSub
     | INT # int
     | ID # id
     | '(' expr ')' # parens
     ;

// ID : identifier，标识符，其实就是变量名
ID : [a-zA-Z]+;
INT : [0-9]+;
NEWLINE : '\r' ? '\n';
WS : [ \t]+ -> skip;
MUL : '*';
DIV : '/';
ADD : '+';
SUB : '-';
