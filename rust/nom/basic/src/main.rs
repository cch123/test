#[macro_use]
extern crate nom;
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

fn main() {
    println!("{:#?}", comp_expr(CompleteStr("aa >= 1")));
    println!("{:#?}", comp_expr(CompleteStr("b = 2")));
    println!("{:#?}", comp_expr(CompleteStr("b in 2")));
}
