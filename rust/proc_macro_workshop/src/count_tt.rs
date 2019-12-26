#[macro_export]
macro_rules! count_tokens {
    ()      => (0usize);
    ($a:tt) => (1usize);
    ($($a:tt $b:tt)*) => (
        count_tokens!($($a)*) * 2
    );
    ($a:tt $($b:tt $c:tt)*) => (
        count_tokens!($($b)*) * 2 + 1
    );
}

/*
这里面的 * 和 + 表示的是“规则”重复多次，和正则类似，如 \d+

但是不一样的是，这里只是说规则重复，不代表规则匹配到的内容本身是一致的

例如：

a b c d

可以认为是 token 规则 * 4

macro 中的 tt 定义：

pub enum TokenTree {
    Group(Group), => 括号括住的内容, A delimited token stream. A Group internally contains a TokenStream which is surrounded by Delimiters.
    Ident(Ident), => 标志符, An identifier (ident).
    Punct(Punct), => An Punct is an single punctuation character like +, - or #. Multi-character operators like += are represented as two instances of Punct with different forms of Spacing returned.
    Literal(Literal), => A literal string ("hello"), byte string (b"hello"), character ('a'), byte character (b'a'), an integer or floating point number with or without a suffix (1, 1u8, 2.3, 2.3f32). Boolean literals like true and false do not belong here, they are Idents.
}
*/