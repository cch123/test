//#[macro_use] // 新版本 rust 没有 macro_use 也行
pub mod easy;
pub mod linkedlist;
pub use linkedlist::Node;

pub mod match_rule;

pub mod bool_expr;
pub mod count_tt;

/*
注意，没有 pub use 这句话，会报错
error[E0433]: failed to resolve: maybe a missing crate `Node`?
  --> src/linkedlist.rs:20:51
   |
20 |             let mut dummy_head = Box::new($crate::Node::new(0));
   |                                                   ^^^^ maybe a missing crate `Node`?
   |
*/

fn main() {
    let x = linkedlist![1,2,3,4,5];
    dbg!(x);

    // macro in easy
    println!("{}", four!());

    // macro in match rule
    /*
    // 内置的 expr 规则是不支持 backtrack 的
    // 如，调用 dead_rule 会报下列错误
    // dead_rule!(x+);
error: expected expression, found end of macro arguments
  --> src/main.rs:41:18
   |
41 |     dead_rule!(x+);
   |                  ^ expected expression

error: aborting due to 2 previous errors
    */

    // 但是裸的 token 匹配没关系
    aplusb!(a+);

    println!("{}", count_tokens!(a , sdfsdf,d df x));
    //count_tokens!(a , sdfsdf,d df x);

}
