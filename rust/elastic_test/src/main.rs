extern crate elastic_query;

fn main() {
    let result = elastic_query::convert("a = 1 and b in (1,2,3)".to_string(), 0, 100, vec![], vec![]).unwrap();
    println!("{}", result);
}
