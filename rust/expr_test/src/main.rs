extern crate pest;
#[macro_use]
extern crate pest_derive;

use pest::Parser;

#[derive(Parser)]
#[grammar = "expr.pest"]
pub struct ExprParser;

fn main() {
    let expr = ExprParser::parse(Rule::expr, "a = 1").expect("parse failed").next().unwrap();
    dbg!(expr);
}
