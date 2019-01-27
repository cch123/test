extern crate serde_json;

use serde_json::json;
fn main() {
    let s = "abc".to_string();
    let a = json!({
        s.clone() : 123
    });
    println!("{}", a);

    let a = json!({
        s.clone() : format!("{}", "fuck")
    });
    println!("{}", a);
}

