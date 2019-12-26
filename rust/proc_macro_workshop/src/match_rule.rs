// https://danielkeep.github.io/tlborm/book/mbe-min-captures-and-expansion-redux.html
/*
Once the parser begins consuming tokens for a capture,
it cannot stop or backtrack.
This means that the second rule of the following macro
cannot ever match, no matter what input is provided:
*/

// 简而言之，不支持 backtrack，所以解析过程是不支持回退的
// 在这种情况下，如果输入一个 dead_rule!(x+)
// 会报错
/*
error: expected expression, found end of macro arguments
  --> src/main.rs:26:18
   |
26 |     dead_rule!(x+);
   |                  ^ expected expression

error: aborting due to previous error
*/
#[macro_export]
macro_rules! dead_rule {
    ($e:expr) => {};
    ($i:ident +) => {};
}

// 下面这种用普通 token 做匹配的还是可以第一条不通继续向下走的
// 和上面的不一样
#[macro_export]
macro_rules! aplusb {
    ( a + b) => {
        println!("match the a + b rule");
    };
    ( a + ) => {
        println!("match the a + rule");
    };
}

/*
if this macro is invoked as dead_rule!(x+).
The interpreter will start at the first rule,
and attempt to parse the input as an expression.
The first token (x) is valid as an expression.
The second token is also valid in an expression,
forming a binary addition node.

At this point, given that there is no right-hand side of the addition,
you might expect the parser to give up and try the next rule. Instead, the parser will panic and abort the entire compilation, citing a syntax error.

As such, it is important in general that
you write macro rules
from most-specific to least-specific.
*/
