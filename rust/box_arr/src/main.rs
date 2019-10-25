#[derive(Debug)]
struct Solution {
    x : Box<[i32]>,
    y : [i32;3],
    //z : [i32], -> println!("{:?}", s);
    //     ^^^^^^^^^^^^^^^^^^^^ doesn't have a size known at compile-time
}

fn main() {
    let s = Solution {
        x : Box::new([1,2]),
        y : [3,2,1],
    };
    println!("{:?}", s);
}
