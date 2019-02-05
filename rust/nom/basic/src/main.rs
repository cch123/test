#[macro_use]
extern crate nom;
use nom::{do_parse, map_res, named, tag, take_while, tag_no_case, multispace};
use nom::types::CompleteStr;

#[derive(Debug, PartialEq)]
pub struct BinOp <'a>{
    pub lhs: CompleteStr<'a>,
    pub op: CompleteStr<'a>,
    pub rhs: CompleteStr<'a>,
}

fn is_ident(c: char) -> bool {
    c.is_ascii_alphanumeric()
}

named!(ident<CompleteStr, CompleteStr>,
  take_while!(is_ident)
);

fn is_op(c:char) -> bool {
    match c {
        '=' | '>' | '<'  => true,
        _ => false
    }
}

named!(operator<CompleteStr, CompleteStr>,
  take_while!(is_op)
);

fn is_space(c:char) -> bool {
    if c == ' ' {
        return true
    }
    false
}
named!(eat_space<CompleteStr, CompleteStr>,
  take_while!(is_space)
);

named!(pub opt_multispace<CompleteStr, Option<CompleteStr>>,
       opt!(complete!(multispace))
);
fn is_digit(c:char) -> bool {
    c.is_ascii_digit()
}

named!(number<CompleteStr, CompleteStr>,
  take_while1!(is_digit)
);

named!(hex_color<CompleteStr, BinOp>,
  do_parse!(
    lhs:   ident >>
    opt_multispace >>
    op:   operator >>
    opt_multispace >>
    rhs:   number>>
    (BinOp{ lhs, op, rhs})
  )
);

fn main() {
    println!("{:#?}", hex_color(CompleteStr("aa = 1")));
    println!("{:#?}", hex_color(CompleteStr("b = 2")));
}
