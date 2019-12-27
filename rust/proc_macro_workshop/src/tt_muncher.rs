// https://danielkeep.github.io/tlborm/book/pat-incremental-tt-munchers.html
macro_rules! mixed_rules {
    () => {};
    (trace $name:ident; $($tail:tt)*) => {
        {
            println!(concat!(stringify!($name), " = {:?}"), $name);
            mixed_rules!($($tail)*);
        }
    };
    (trace $name:ident = $init:expr; $($tail:tt)*) => {
        {
            let $name = $init;
            println!(concat!(stringify!($name), " = {:?}"), $name);
            mixed_rules!($($tail)*);
        }
    };
}

pub fn tt_munchers() {
    let (a, b, c, d, x) = (1, 2, 3, 4, 5);
    mixed_rules!(trace a;trace b = 3;trace  d= 4;trace x;);
}
