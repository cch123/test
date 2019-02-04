grammar ArrayInit;
// grammar 是语法文件开头，ArrayInit 这个名字必须与文件名相匹配
// 典型的 java style
// 下面这些是语法分析器
init : '{' value (',' value)* '}'; // 至少匹配一个 value，antlr 里没有 + 的语法？不应该吧

// 表示一个 value 可以是嵌套花括号，也可以是一个简单数值，用 INT 词法符号表示
value : init
    | INT
    ;

// 语法分析器的规则必须是小写字母开头，词法分析器的规则必须用大写字母开头
// 下面这些是词法分析器
INT : [0-9]+; // 定义词法符号 INT，和正则表达式差不多
WS : [ \t\r\n]+ -> skip; // 这条规则的意思是空白字符直接丢弃
