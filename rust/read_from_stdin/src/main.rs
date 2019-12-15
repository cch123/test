use std::io::{self, Read};

fn main() ->  io::Result<()> {
    let mut buffer = Vec::new();
    io::stdin().read_to_end(&mut buffer)?;
    println!("read result : {:?}", String::from_utf8(buffer));
    Ok(())
}
