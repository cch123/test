extern crate combine;
use combine::char::char;
use combine::{Parser, many1};

fn main() {
    let r = many1::<String, _>(char('a').or(char('b'))).parse("aaaabbbbabababab");
    match r {
        Ok((value, remaining)) => println!("value: {} remaining: {}", value,remaining),
        Err(err) => println!("{}", err)
    }
}
