grammar where;

prog: expr;

expr: expr 'and' expr # andExpr
| expr 'or' expr # orExpr
| '(' expr ')' # quotedExpr
| leafExpr # cmpExpr
;

leafExpr: ID op='=' val
| ID op='>' val
| ID op='>=' val
| ID op='<' val
| ID op='<=' val
| ID op='in' val
| ID op='not in' val
;

val: STRING | INT | FLOAT | BOOL;

ID : [a-zA-Z]+;
STRING: '\'' .*? '\'';
FLOAT: [0-9]+'.'[0-9]*;
INT: [0-9]+;
BOOL: 'true' | 'false';
WS : [ \t] + -> skip;
