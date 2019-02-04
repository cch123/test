lexer grammar CommonLexerRules; // 注意开头的 lexer grammar

ID : [a-zA-Z]+;
INT : [0-9]+;
NEWLINE:'\r'?'\n';
WS : [ \t]+ -> skip; // 丢弃空白字符
