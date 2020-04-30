// https://blog.rust-lang.org/2020/03/12/Rust-1.42.html
fn foo(words: &[&str]) {
    match words {
        ["Hello", "World", "!", ..] => println!("Hello World!"),
        ["Foo", "Bar", ..] => println!("Baz"),
        /*
        The .. is called a "rest pattern," 
        because it matches the rest of the slice. 
        The above example uses the rest pattern
        at the end of a slice, but you can also
        use it in other ways:
        */
        rest => println!("{:?}", rest),
    }
}

fn main() {
    let arr = ["Hello", "World", "!"];
    foo(&arr);
    let arr = ["ffff"];
    foo(&arr);
}

fn foo2(words: &[&str]) {
    match words {
        // Ignore everything but the last element, which must be "!".
        [.., "!"] => println!("!!!"),

        // `start` is a slice of everything except the last element, which must be "z".
        [start @ .., "z"] => println!("starts with: {:?}", start),

        // `end` is a slice of everything but the first element, which must be "a".
        ["a", end @ ..] => println!("ends with: {:?}", end),

        rest => println!("{:?}", rest),
    }
}