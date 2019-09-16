#[macro_use]
extern crate nom;
use nom::alt;
use nom::fold_many0;
use nom::types::CompleteStr;
use nom::{do_parse, multispace, named, take_while};

#[derive(Debug, PartialEq)]
pub enum BoolExpr<'a> {
    CompExpr {
        lhs: CompleteStr<'a>,
        op: CompleteStr<'a>,
        rhs: CompleteStr<'a>,
    },
    AndExpr {
        left: Box<BoolExpr<'a>>,
        right: Box<BoolExpr<'a>>,
    },
    OrExpr {
        left: Box<BoolExpr<'a>>,
        right: Box<BoolExpr<'a>>,
    },
}

named!(ident<CompleteStr, CompleteStr>,
  take_while!(|c: char| c.is_ascii_alphanumeric())
);

named!(operator<CompleteStr, CompleteStr>,
  take_while!(|c: char| {
    if c.is_ascii_alphanumeric() {return true}
    match c{ '='|'>'|'<' => true, _ => false}
  })
);

named!(eat_space<CompleteStr, CompleteStr>,
  take_while!(|c:char| c.is_whitespace())
);

named!(pub opt_multispace<CompleteStr, Option<CompleteStr>>,
  opt!(complete!(multispace))
);

named!(number<CompleteStr, CompleteStr>,
  take_while1!(|c : char| c.is_ascii_digit())
);

named!(comp_expr<CompleteStr, BoolExpr>,
  do_parse!(
    lhs:   ident >> opt_multispace >>
    op:   operator >> opt_multispace >>
    rhs:   number>>
    (BoolExpr::CompExpr{ lhs, op, rhs})
  )
);

named!(or_expr<CompleteStr, BoolExpr>,
  do_parse!(
    left: comp_expr >> opt_multispace >>
    tag_no_case!("or") >> opt_multispace >>
    right : bool_expr >>
    (BoolExpr::OrExpr{left: Box::new(left), right: Box::new(right)})
  )
);

named!(and_expr<CompleteStr, BoolExpr>,
  do_parse!(
    left: comp_expr >> opt_multispace >>
    tag_no_case!("and") >> opt_multispace >>
    right : bool_expr >>
    (BoolExpr::AndExpr{left: Box::new(left), right: Box::new(right)})
  )
);

named!(bool_expr<CompleteStr, BoolExpr>,
  alt!(
    and_expr | or_expr | comp_expr
  )
);

fn main() {
    println!("{:#?}", bool_expr(CompleteStr("aa >= 1")));
    println!("{:#?}", bool_expr(CompleteStr("b = 2")));
    println!("{:#?}", bool_expr(CompleteStr("b in 2")));
    println!("{:#?}", bool_expr(CompleteStr("b in 2 and a = 1")));
    println!(
        "{:#?}",
        bool_expr(CompleteStr("b in 2 and a = 1 and c = 2"))
    );
    (1..100).for_each(|x| {});

    // TODO fix this
    //println!("{:#?}", bool_expr(CompleteStr("b in 2 and a = 1 and c =2")));
}

impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        let mut res = vec![0, 1, 1];
        (3..=n).for_each(|i| {
            let mut i = i as usize;
            res.push(res.get(i - 1).unwrap() + res.get(i - 2).unwrap());
        });
        res[n as usize]
    }
}
struct Solution;

impl Solution {
    pub fn day_of_year(date: String) -> i32 {
        let day_for_normal = vec![31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        let day_for_leap = vec![31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        let ymd: Vec<&str> = date.split("-").collect();
        let (y, m, d) = (
            ymd[0].parse::<i32>().unwrap(),
            ymd[1].parse::<i32>().unwrap(),
            ymd[2].parse::<i32>().unwrap(),
        );
        let mut should_use = day_for_normal;
        if Self::is_leap(y) {
            should_use = day_for_leap;
        }
        (0..m - 1).map(|i| should_use[i as usize]).sum() + d
    }
    fn is_leap(y: i32) -> bool {
        (y % 4 == 0 && y % 100 != 0) || y % 400 == 0
    }
}
