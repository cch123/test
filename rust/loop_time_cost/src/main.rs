use std::time::Instant;
use std::collections::LinkedList;

fn main() {
    let start = Instant::now();
    let mut v = Vec::with_capacity(0);
    (0..1000000).for_each(|i|{
         // do nothing
         v.push(i);
    });
    dbg!(start.elapsed());
    dbg!(v.len());

    let start = Instant::now();
    let mut list = LinkedList::new();
    (0..1000000).for_each(|i|{
         // do nothing
         //list.push(i);
         list.push_back(i);
    });
    dbg!(start.elapsed());
    dbg!(list.len());
}
