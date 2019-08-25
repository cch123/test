#[macro_use]
extern crate lazy_static;
use std::collections::HashMap;

lazy_static! {
    static ref m: HashMap<i32, &'static str> =
        { vec![(1, "abc"), (2, "def")].into_iter().collect() };
}

lazy_static! {
    static ref HASHMAP: HashMap<u32, &'static str> = {
        let mut ma = HashMap::new();
        ma.insert(0, "foo");
        ma.insert(1, "bar");
        ma.insert(2, "baz");
        ma
    };
}

fn main() {
    println!("{:?}", m.get(&1));
}
