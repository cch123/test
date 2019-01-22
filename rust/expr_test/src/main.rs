/*
基本的语法设计，可以参考：

https://github.com/sejr/iron/blob/master/grammar/iron.pest

如果是写 bool expr 的解析的话，这里面的范例应该是已经够了
*/

extern crate pest;
#[macro_use]
extern crate pest_derive;

use pest::Parser;

#[derive(Parser)]
#[grammar = "expr.pest"]
pub struct ExprParser;

fn main() {
    /*
    let expr = ExprParser::parse(Rule::expr, r#"a = "2121""#).expect("parse failed").next().unwrap();
    dbg!(expr);
    let expr = ExprParser::parse(Rule::expr, r#"a = 1 and b = 2"#).expect("parse failed").next().unwrap();
    dbg!(expr);
    let expr = ExprParser::parse(Rule::expr, "(a=1) and (b=2)").expect("parse failed").next().unwrap();
    dbg!(expr);
    let expr = ExprParser::parse(Rule::expr, "a=1 and b=2").expect("parse failed").next().unwrap();
    dbg!(expr);
    */
    let expr = ExprParser::parse(Rule::expr, "a=1 and ((b = 2) and c=1)").expect("parse failed").next().unwrap();
    dbg!(expr);
}

/*
[src/main.rs:13] expr = Pair {
    rule: expr,
    span: Span {
        str: "a = \"1\"",
        start: 0,
        end: 7
    },
    inner: [
        Pair {
            rule: binary_op,
            span: Span {
                str: "a = \"1\"",
                start: 0,
                end: 7
            },
            inner: [
                Pair {
                    rule: field,
                    span: Span {
                        str: "a ",
                        start: 0,
                        end: 2
                    },
                    inner: []
                },
                Pair {
                    rule: op,
                    span: Span {
                        str: "=",
                        start: 2,
                        end: 3
                    },
                    inner: []
                },
                Pair {
                    rule: value,
                    span: Span {
                        str: "\"1\"",
                        start: 4,
                        end: 7
                    },
                    inner: [
                        Pair {
                            rule: string_literal,
                            span: Span {
                                str: "\"1\"",
                                start: 4,
                                end: 7
                            },
                            inner: [
                                Pair {
                                    rule: string,
                                    span: Span {
                                        str: "1",
                                        start: 5,
                                        end: 6
                                    },
                                    inner: []
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}

*/