use std::time::{Duration, Instant};

fn main() {
    let start = Instant::now();
    (0..1000000).for_each(|_|{
         // do nothing
    });
    dbg!(start.elapsed());
}
