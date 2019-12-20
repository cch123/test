use a::Person;
use a::New;
fn main() {
    /*
    error[E0639]: cannot create non-exhaustive struct using struct expression
 --> src/main.rs:3:13
  |
3 |     let p = Person {age :1};
  |
    */
    dbg!(New());
}
