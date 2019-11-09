use async_std::{fs::File, prelude::*, task, };

fn main() {
    let n = task::block_on(async {
        let mut file = File::open("./Cargo.toml").await?;
        let mut contents = String::new();
        file.read_to_string(&mut contents).await?;
        /*
        Unfortunately, there's currently no way to "give fut a type",
        nor a way to explicitly specify the return type of an async block.
        To work around this, use the "turbofish" operator
        to supply the success and error types for the async block:
        */
        // async block turbofish
        // https://rust-lang.github.io/async-book/07_workarounds/03_err_in_async_blocks.html
        Ok::<String, Box<dyn std::error::Error>>(contents)
    });
    println!("{:?}", n);
}
