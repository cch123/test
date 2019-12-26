/*
follow 规则：
item: anything.
block: anything.
stmt: => , ;
pat: => , = if in
expr: => , ;
ty: , => : = > ; as
ident: anything.
path: , => : = > ; as
meta: anything.
tt: anything.
*/

#[macro_export]
macro_rules! and_expr {
    () => {
        false
    };
    ($left:expr ; and $right:expr) => {
        if $left && $right {
            true
        } else {
            false
        }
    };
}

#[macro_export]
macro_rules! or_expr {
    () => {
        false
    };
    ($left:expr ; or $right:expr) => {
        if $left || $right {
            true
        } else {
            false
        }
    };
}
