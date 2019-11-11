//! Hello world example for Rust.
use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>>{
    let addr = "0.0.0.0:11333";
    let mut listener = TcpListener::bind(&addr).await?;
    loop {
        let (mut conn, remote) = listener.accept().await?;
        dbg!(remote);

        tokio::spawn(async move {
            let mut buf : [u8;10]= [0; 10];
            loop {
                let n = conn.read(&mut buf).await.expect("expected data, but EOF");
                
                if n == 0 {
                    break;
                }
                println!("received packet, len {}, content : {:?}", n, String::from_utf8(buf[0..n].to_vec()));

                conn.write_all(&buf[0..n]).await.expect("write faile");
                // 如果收到 quit，关闭连接
                //if &buf[0..n] == b"quit\r\n" {
                if buf[0..n].starts_with(b"quit") {
                    conn.shutdown(std::net::Shutdown::Both).expect("shutdown fuck!");
                }
            }
        });
    }
}
