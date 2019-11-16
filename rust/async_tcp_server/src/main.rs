//! Hello world example for Rust.
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpListener;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:11333";
    let mut listener = TcpListener::bind(&addr).await?;
    loop {
        let (mut conn, remote) = listener.accept().await?;
        dbg!(remote);

        tokio::spawn(async move {
            let mut buf = [0; 1024];
            loop {
                let n = conn.read(&mut buf).await.expect("expected data, but EOF");

                if n == 0 {
                    break;
                }
                println!(
                    "received packet, len {}, content : {:?}",
                    n,
                    String::from_utf8(buf[0..n].to_vec())
                );

                conn.write_all(&buf[0..n]).await.expect("write faile");
                // 如果收到 quit，关闭连接
                //if &buf[0..n] == b"quit\r\n" {
                if buf[0..n].starts_with(b"quit") {
                    conn.shutdown(std::net::Shutdown::Both)
                        .expect("shutdown fuck!");
                }

                let num_bytes = buf[0..n]
                    .iter()
                    .map(|e| *e)
                    .take_while(|&e| e > b'0' && e < b'9')
                    .collect::<Vec<_>>();

                //println!("string from request: {:?}", String::from_utf8(num_bytes));

                match num_bytes.len() {
                    e if e > 512 => {
                        conn.shutdown(std::net::Shutdown::Both)
                            .expect("shutdown because req too long");
                    }
                    _ => {
                        let req_len = std::str::from_utf8(&num_bytes)
                            .unwrap()
                            .parse::<i32>()
                            .unwrap();
                        let req = String::from_utf8(buf[num_bytes.len() + 1..n].to_vec()).unwrap();
                        if req.trim().len() != req_len as usize {
                            println!("wrong len of request body");
                        }
                    }
                }
            }
        });
    }
}
