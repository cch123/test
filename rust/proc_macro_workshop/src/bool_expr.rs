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
/*
macro_rules! term {
    () => { false };
    ($e:expr) => {
        println!("{} or {}", $left, $right);
    }
}

macro_rules! and_expr {
    () => { false };
    ($left:term or $right)
}
*/