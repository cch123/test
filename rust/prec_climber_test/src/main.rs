extern crate pest;

#[macro_use]
extern crate pest_derive;

use pest::error::Error;

use pest::Parser;
use pest::iterators::Pair;
use pest::prec_climber::{Assoc, Operator, PrecClimber};

fn main() {
   // let input = std::env::args().nth(1).expect("args[1] is not found");
    let input = "1+2+3";
    match do_parse(&input) {
        Ok(expression) => println!("{:#?}", expression),
        Err(err) => println!("failed"),
    }
}



#[derive(Debug)]
enum Expression {
    Number(i32),
    Infix(InfixOp, Box<Expression>, Box<Expression>),
}

impl Expression {
    fn infix<L, R>(op: InfixOp, lhs: L, rhs: R) -> Self
        where
            L: Into<Expression>,
            R: Into<Expression>,
    {
        Expression::Infix(op.into(), Box::new(lhs.into()), Box::new(rhs.into()))
    }
}

#[derive(Debug)]
enum InfixOp {
    Plus,
    Minus,
    Times,
    Divide,
    Modulus,
    Power,
}



const _GRAMMAR: &str = include_str!("grammar.pest");

#[derive(Parser)]
#[grammar = "grammar.pest"]
struct ExprParser;

fn do_parse<'s>(input: &'s str) -> Result<Expression, Error<String>> {
    let mut pairs = ExprParser::parse(Rule::expression, input).expect("fuck");
    let pair = pairs.next().unwrap();
    Ok(into_expression(pair))
}

fn into_expression(pair: Pair<Rule>) -> Expression {
    let climber = PrecClimber::new(vec![
        Operator::new(Rule::plus, Assoc::Left) |
            Operator::new(Rule::minus, Assoc::Left),
        Operator::new(Rule::times, Assoc::Left) | Operator::new(Rule::divide, Assoc::Left) |
            Operator::new(Rule::modulus, Assoc::Left),
        Operator::new(Rule::power, Assoc::Right),
    ]);

    consume(pair, &climber)
}

fn consume(pair: Pair<Rule>, climber: &PrecClimber<Rule>) -> Expression {
    // println!("Rule: {:?}", pair.as_rule());
    // println!("Text: {:?}", pair.as_str());
    // println!();

    let primary = |pair| consume(pair, climber);
    let infix = |lhs, op: Pair<Rule>, rhs| match op.as_rule() {
        Rule::plus => Expression::infix(InfixOp::Plus, lhs, rhs),
        Rule::minus => Expression::infix(InfixOp::Minus, lhs, rhs),
        Rule::times => Expression::infix(InfixOp::Times, lhs, rhs),
        Rule::divide => Expression::infix(InfixOp::Divide, lhs, rhs),
        Rule::modulus => Expression::infix(InfixOp::Modulus, lhs, rhs),
        Rule::power => Expression::infix(InfixOp::Power, lhs, rhs),
        _ => unreachable!(),
    };
    match pair.as_rule() {
        Rule::expression => {
            let pairs = pair.into_inner();
            climber.climb(pairs, primary, infix)
        }
        Rule::primary => pair.into_inner().next().map(primary).unwrap(),
        Rule::number => {
            let number = pair.as_str().parse().unwrap();
            Expression::Number(number)
        }
        _ => unreachable!(),
    }
}