macro_rules! call_with_larch {
    ($callback:ident) => {
        $callback!(larch)
    };
}

macro_rules! expand_to_larch {
    () => {
        larch
    };
}

macro_rules! recognise_tree {
    (larch) => {
        println!("#1, the Larch.")
    };
    (redwood) => {
        println!("#2, the Mighty Redwood.")
    };
    (fir) => {
        println!("#3, the Fir.")
    };
    (chestnut) => {
        println!("#4, the Horse Chestnut.")
    };
    (pine) => {
        println!("#5, the Scots Pine.")
    };
    ($($other:tt)*) => {
        println!("I don't know; some kind of birch maybe?")
    };
}

macro_rules! callback {
    ($callback:ident($($args:tt)*)) => {
        $callback!($($args)*)
    };
}
/*
Due to the order that macros are expanded in, it is (as of Rust 1.2) impossible to pass information to a macro from the expansion of another macro. This can make modularising macros very difficult.
An alternative is to use recursion and pass a callback. Here is a trace of the above example to demonstrate how this takes place:
*/
// 这个函数是想证明，宏是从外向内扩展的
// 所以，通过一个宏展开一些内容，再把这些内容传给另一个宏是没有办法实现的
// 这种情况下就只能传一个 callback 给宏来实现类似的需求了
pub fn callback() {
    // 是先展开 recognise_tree
    // 括号内的内容传进去之后，被当成 tt 来匹配了
    recognise_tree!(expand_to_larch!());

    // 把 recognise_tree 当成一个名字传进 call_with_larch 中
    call_with_larch!(recognise_tree);

    // 调用宏，需要保证宏在该函数之前定义
    // 要不会报未找到，和 c 有点类似
    callback!(callback(println("Yes, this *was* unnecessary.")));
}

