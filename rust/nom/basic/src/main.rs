#[macro_use]
extern crate nom;
use nom::{do_parse, map_res, named, tag, take_while, tag_no_case, multispace};

#[derive(Debug, PartialEq)]
pub struct BinOp <'a>{
    pub lhs: &'a str,
    pub op: &'a str,
    pub rhs: &'a str,
}

fn is_ident(c: char) -> bool {
    c.is_ascii_alphanumeric()
}

named!(ident<&str, &str>,
  take_while!(is_ident)
);

fn is_op(c:char) -> bool {
    match c {
        '=' | '>' | '<'  => true,
        _ => false
    }
}

named!(operator<&str, &str>,
  take_while!(is_op)
);

fn is_space(c:char) -> bool {
    if c == ' ' {
        return true
    }
    false
}
named!(eat_space<&str, &str>,
  take_while!(is_space)
);

named!(pub opt_multispace<&str, Option<&str>>,
       opt!(complete!(multispace))
);
fn is_digit(c:char) -> bool {
    c.is_ascii_digit()
}

named!(number<&str, &str>,
  take_while1!(is_digit)
);

named!(hex_color<&str, BinOp>,
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
    println!("{:#?}", hex_color("aa = 1 "));
    println!("{:#?}", hex_color("b = 2 "));
}
