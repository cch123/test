use std::marker::PhantomData;
// compile error
//  rustc --explain E0392
/*
T 指向的内容有生命周期
但是这种这里的 raw pointer 没有这个东西
struct Slice<'a, T> {
    start: *const T,
    end: *const T,
}
*/
struct Slicex<T> {
    start: *const T,
}

// 生命周期参数在 raw pointer 没体现出来
// 所以额外加一个 phantomdata
struct Slice<'a, T> {
    start: *const T,
    end: *const T,
    phantom : PhantomData<&'a T>,
}

fn main() {
}

// https://www.jianshu.com/p/8554bbf13a02

// 这篇里的例子有点莫名其妙
// https://doc.rust-lang.org/rust-by-example/generics/phantom.html