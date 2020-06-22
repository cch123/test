//! Hello world example for Rust.
use tokio::prelude::*;
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
                    .take_while(|&e| e >= b'0' && e <= b'9')
                    .collect::<Vec<_>>();

                //println!("string from request: {:?}", String::from_utf8(num_bytes));

                match num_bytes.len() {
                    l if l > 20 => {
                        println!("shutdown because request too long");
                        conn.shutdown(std::net::Shutdown::Both)
                            .expect("shutdown failed");
                    }
                    l if l == 0 => {
                        println!("shutdown because no request length description");
                        conn.shutdown(std::net::Shutdown::Both)
                            .expect("shutdown failed");
                    }
                    l => {
                        let req_len = String::from_utf8(num_bytes)
                            .unwrap()
                            .parse::<i64>()
                            .unwrap();

                        let buf_vec = buf[l+1..n].to_vec();
                        let req_body = std::str::from_utf8(&buf_vec).unwrap();

                        let mut req_body = req_body.trim_start().to_string();

                        while req_body.len() < (req_len as usize) {
                            let n = conn.read(&mut buf).await.expect("expected data, but EOF");
                            println!("fyckk");

                            if n == 0 {
                                println!("fyc");
                                break;
                            }
                            req_body.push_str(std::str::from_utf8(&buf[..n].to_vec()).unwrap());
                        }
                        req_body = req_body.trim_end().to_string();

                        if req_body.len() != (req_len as usize) {
                            println!("wrong len of request body");
                        }

                        dbg!(req_len, req_body.clone(), req_body.len());
                    }
                }
            }
        });
    }
}
