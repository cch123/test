#[macro_use]
extern crate nom;
use nom::types::CompleteStr;
use nom::{do_parse, multispace, named, take_while};
use nom::alt_complete;

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
  alt_complete!(
    and_expr | or_expr | comp_expr
  )
);

fn main() {
    println!("{:#?}", bool_expr(CompleteStr("aa >= 1")));
    println!("{:#?}", bool_expr(CompleteStr("b = 2")));
    println!("{:#?}", bool_expr(CompleteStr("b in 2")));
    println!("{:#?}", bool_expr(CompleteStr("b in 2 and a = 1")));
    println!("{:#?}", bool_expr(CompleteStr("b in 2 and a = 1 and c = 2")));

    // TODO fix this
    //println!("{:#?}", bool_expr(CompleteStr("b in 2 and a = 1 and c =2")));
}
