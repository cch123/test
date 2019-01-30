extern crate tokio;
use tokio::io::copy;
use tokio::net::TcpListener;
use tokio::prelude::*;

fn main() {
    let addr = "127.0.0.1:12345".parse().unwrap();
    let listener = TcpListener::bind(&addr).expect("unable to bind tcp addr");

    let server = listener
        .incoming()
        .map_err(|e| eprintln!("accept failed = {:?}", e))
        .for_each(|sock| {
            let (reader, writer) = sock.split();
            let bytes_copied = copy(reader, writer);
            let handle_conn = bytes_copied
                .map(|amt| println!("wrote {:} bytes", amt.0))
                .map_err(|err| eprintln!("IO error {:?}", err));

            tokio::spawn(handle_conn)
        });

    tokio::run(server);
}

