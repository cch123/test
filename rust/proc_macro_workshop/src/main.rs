#![feature(trace_macros)]
//#[macro_use] // 新版本 rust 没有 macro_use 也行
pub mod easy;
pub mod linkedlist;
pub use linkedlist::Node;

pub mod match_rule;

pub mod bool_expr;
pub mod count_tt;
pub mod debug;
pub mod substitution_is_not_token_based;

pub mod callback;

pub mod tt_muncher;
pub mod internal_rule;

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
    let x = linkedlist![1, 2, 3, 4, 5];
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

    println!("count tokens macro: {}", count_tokens!(a , sdfsdf,d df x));

    let (a, b) = (10, -1);
    println!("{}", and_expr!( a > 1; and b <0 ));
    println!("{}", or_expr!( a > 1; or b <0 ));

    // stringify!(tt), stringifying a sequence of token trees
    println!("{:?}", stringify!(dummy(2 * (1 + (3)))));
    // stringify!(expr), stringifying an AST expression node.
    println!("{:?}", capture_expr_then_stringify!(dummy(2 * (1 + (3)))));

    // -------------------------- 对比一下这个和下面那个的输出结果
    println!(
        "{}\n{}\n{}\n{}",
        what_is!(#[no_mangle]),
        what_is!(#[inline]),
        capture_then_what_is!(#[no_mangle]),
        capture_then_what_is!(#[inline]),
    );

    println!(
        "{}\n{}\n{}\n{}",
        what_is!(#[no_mangle]),
        what_is!(#[inline]),
        capture_then_what_is2!(#[no_mangle]),
        capture_then_what_is2!(#[inline]),
    );
    // ----------------------------------------------------------

    each_tt!(foo bar baz quux);
    trace_macros!(true);
    /*
          --> src/main.rs:79:5
       |
    79 |     each_tt!(spim wak plee whum);
       |     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
       |
       = note: expanding `each_tt! { spim wak plee whum }`
       = note: to `each_tt ! (wak plee whum) ;`
       = note: expanding `each_tt! { wak plee whum }`
       = note: to `each_tt ! (plee whum) ;`
       = note: expanding `each_tt! { plee whum }`
       = note: to `each_tt ! (whum) ;`
       = note: expanding `each_tt! { whum }`
       = note: to `each_tt ! () ;`
       = note: expanding `each_tt! {  }`
       = note: to ``
        */
    each_tt!(spim wak plee whum);
    trace_macros!(false);
    each_tt!(trom qlip winp xod);

    callback::callback();

    tt_muncher::tt_munchers();

    internal_rule::internal_rule();
    crate_name_util!(@count_tts);
}
