/*
基本的语法设计，可以参考：

https://github.com/sejr/iron/blob/master/grammar/iron.pest

如果是写 bool expr 的解析的话，这里面的范例应该是已经够了
*/

extern crate pest;
#[macro_use]
extern crate pest_derive;

use pest::error::Error;
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
    let expr = ExprParser::parse(Rule::expr, "a=1 and ((b = 2) and c=1)").expect("parse failed").next().unwrap();
    dbg!(expr);
    let expr = ExprParser::parse(Rule::expr, "a in 1").expect("parse failed").next().unwrap();
    dbg!(expr);
    */
    let expr = ExprParser::parse(Rule::expr, "(a=1 and ((b = 2) and c=1))")
        .expect("parse failed")
        .next()
        .unwrap();
    //dbg!(&expr);
    parse_expr(expr).unwrap();
}

use pest::iterators::Pair;

enum Node {
    AndExpr {
        left: Box<Node>,
        right: Box<Node>,
    },
    OrExpr {
        left: Box<Node>,
        right: Box<Node>,
    },
    CompExpr {
        lhs: String,
        op: String,
        rhs: String,
    },
    Null,
}

fn parse_expr(expr: Pair<Rule>) -> Result<Node, Error<Rule>> {
    for record in expr.into_inner() {
        match record.as_rule() {
            Rule::expr => {
                record.into_inner().for_each(|r| {
                    parse_expr(r).unwrap();
                });
            },
            Rule::or_expr => {}
            Rule::and_expr => {
                let mut iter = record.into_inner();
                let left = iter.next().unwrap();
                let right = iter.next().unwrap();
                let left_tree = parse_expr(left).unwrap();
                let right_tree= parse_expr(right).unwrap();
                return Ok(Node::AndExpr {
                    left : Box::new(left_tree),
                    right: Box::new(right_tree),
                });
            },
            Rule::paren_bool => {
                return parse_expr(record.into_inner().next().unwrap());
            }
            Rule::comp_expr => {
                return Ok(Node::CompExpr {
                    lhs: "abc".to_string(),
                    op: "=".to_string(),
                    rhs: "aaa".to_string(),
                })
            },
            _ => {
                println!("fff{:?}", record);
            }
        }
    }
    Ok(Node::Null)
}
