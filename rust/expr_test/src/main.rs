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
    let expr = ExprParser::parse(Rule::expr, r#"a = "2121""#)
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, r#"a = 1 and b = 2"#)
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, "(a=1) and (b=2)")
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, "a=1 and b=2")
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, "a=1 and ((b = 2) and c=1)")
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, "a in 1")
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
    let expr = ExprParser::parse(Rule::expr, "(a=1 and ((b = 2) and c=1))")
        .expect("parse failed")
        .next()
        .unwrap();
    let tree = parse_expr(expr).unwrap();
    dbg!(tree);
}

use pest::iterators::Pair;

#[derive(Debug)]
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
}

fn parse_expr(record: Pair<Rule>) -> Result<Node, Error<Rule>> {
    match record.clone().as_rule() {
        Rule::expr => {
            return parse_expr(record.into_inner().next().unwrap());
        }
        Rule::or_expr => {
            let mut iter = record.into_inner();
            let (left_tree, right_tree) = (
                parse_expr(iter.next().unwrap()).unwrap(),
                parse_expr(iter.next().unwrap()).unwrap(),
            );
            return Ok(Node::OrExpr {
                left: Box::new(left_tree),
                right: Box::new(right_tree),
            });
        }
        Rule::and_expr => {
            let mut iter = record.into_inner();
            let (left_tree, right_tree) = (
                parse_expr(iter.next().unwrap()).unwrap(),
                parse_expr(iter.next().unwrap()).unwrap(),
            );
            return Ok(Node::AndExpr {
                left: Box::new(left_tree),
                right: Box::new(right_tree),
            });
        }
        Rule::paren_bool => {
            return parse_expr(record.into_inner().next().unwrap());
        }
        Rule::comp_expr => {
            let mut iter = record.into_inner();
            let (field, op, value) = (
                iter.next().unwrap().as_str().to_string(),
                iter.next().unwrap().as_str().to_string(),
                iter.next().unwrap().as_str().to_string(),
            );

            return Ok(Node::CompExpr {
                lhs: field,
                op,
                rhs: value,
            });
        }
        _ => unreachable!(),
    }
}
