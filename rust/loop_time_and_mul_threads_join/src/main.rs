use std::collections::LinkedList;
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Instant;

fn multi_thread_push() {
    let v = Arc::new(Mutex::new(Vec::new()));
    let mut threads = vec![];
    for i in 0..10 {
        let v = v.clone();
        let t = thread::spawn(move || {
            v.lock().unwrap().push(i);
        });
        threads.push(t);
    }
    for t in threads {
        t.join().unwrap();
    }
    dbg!(v);
}

fn main() {
    multi_thread_push();

    let start = Instant::now();
    let mut v = Vec::with_capacity(0);
    (0..1000000).for_each(|i| {
        // do nothing
        v.push(i);
    });
    dbg!(start.elapsed());
    dbg!(v.len());

    let start = Instant::now();
    let mut list = LinkedList::new();
    (0..1000000).for_each(|i| {
        // do nothing
        //list.push(i);
        list.push_back(i);
    });
    dbg!(start.elapsed());
    dbg!(list.len());
}
