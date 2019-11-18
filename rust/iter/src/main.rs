fn main() {
    let mut v = vec![1,2,3];
    v.iter().for_each(|e| {
        println!("{}", e);
    });

    (0..10).for_each(|e| {
        println!("{}", e);
    });

    println!("{:?}", v.iter().map(|e| e + 1).collect::<Vec<_>>());
    v.iter_mut().for_each(|e| {*e = 1});
    dbg!(v);
}
