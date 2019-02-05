#[macro_use]
extern crate nom;
use nom::{do_parse, named, take_while, multispace};
use nom::types::CompleteStr;

#[derive(Debug, PartialEq)]
pub struct CompExpr <'a>{
    pub lhs: CompleteStr<'a>,
    pub op: CompleteStr<'a>,
    pub rhs: CompleteStr<'a>,
}

named!(ident<CompleteStr, CompleteStr>,
  take_while!(|c: char| c.is_ascii_alphanumeric())
);


named!(operator<CompleteStr, CompleteStr>,
  take_while!(|c: char| match c{ '='|'>'|'<' => true, _ => false})
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

named!(comp_expr<CompleteStr, CompExpr>,
  do_parse!(
    lhs:   ident >>
    opt_multispace >>
    op:   operator >>
    opt_multispace >>
    rhs:   number>>
    (CompExpr{ lhs, op, rhs})
  )
);

fn main() {
    println!("{:#?}", comp_expr(CompleteStr("aa =1 ")));
    println!("{:#?}", comp_expr(CompleteStr("b = 2")));
}
