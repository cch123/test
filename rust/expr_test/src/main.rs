/*
基本的语法设计，可以参考：

https://github.com/sejr/iron/blob/master/grammar/iron.pest

如果是写 bool expr 的解析的话，这里面的范例应该是已经够了
*/

extern crate pest;
#[macro_use]
extern crate pest_derive;

use pest::Parser;
use pest::error::Error;

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
    let expr = ExprParser::parse(Rule::expr, "(a=1 and ((b = 2) and c=1))").expect("parse failed").next().unwrap();
    //dbg!(&expr);
    parse_expr(expr).unwrap();
}

use pest::iterators::Pair;
fn parse_expr(expr : Pair<Rule>) -> Result<(), Error<Rule>> {
    for record in expr.into_inner() {
        match record.as_rule() {
            Rule::expr => {
                record.into_inner().for_each(|r|{
                    parse_expr(r);
                });
            },
            Rule::or_expr => {
                dbg!(record);
            },
            Rule::and_expr => {
                dbg!(record);
            },
            Rule::paren_bool => {
                record.into_inner().for_each(|r|{
                    dbg!(r);
                });
            }
            Rule::comp_expr => {
            }
            _ => unreachable!()
        }
    }
    Ok(())
}
