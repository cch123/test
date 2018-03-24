use std::io::Write;
use std::str::FromStr;

fn add(n: u64, m: u64) -> u64 {
    return m + n
}

fn main() {
    let mut number_list = Vec::new();

    for arg in std::env::args().skip(1) {
        number_list.push(u64::from_str(&arg)
        .expect("error parsing argument"));
    }

    if number_list.len() == 0 {
        writeln!(std::io::stderr(), "Usage xxx a b").unwrap();
        std::process::exit(1);
    }

    let mut d  = number_list[0];
    for m in &number_list[1..] {
        d = add(d, *m);
    }

    println!("The sum of of {:?} is {}", number_list, d);
}
