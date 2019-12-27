# follow 规则
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

macro_rules! 一般禁止一个  repetition  紧跟另一个 repetition，即使内容本身没有冲突。

另一个可能令人惊奇意外的点是，基于宏的替换并不是基于 token 的，尽管看起来很像。

下面是一个例子：

```rust
macro_rules! capture_expr_then_stringify {
    ($e:expr) => {
        stringify!($e)
    };
}

fn main() {
    println!("{:?}", stringify!(dummy(2 * (1 + (3)))));
    println!("{:?}", capture_expr_then_stringify!(dummy(2 * (1 + (3)))));
}
```

# scope

https://danielkeep.github.io/tlborm/book/mbe-min-scoping.html