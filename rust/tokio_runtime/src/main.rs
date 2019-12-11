use tokio::net::TcpStream;
use tokio::time::delay_for;
use tokio::prelude::*;
use tokio::runtime::Runtime;
use tokio::runtime;

use std::time::Duration;
use std::sync::{
    Arc,
    atomic::{AtomicBool, Ordering}
};
use std::sync::atomic::AtomicI32;

use futures::future::join_all;

/*
async fn connect(addr: &str) -> io::Result<TcpStream> {
    match TcpStream::connect(addr).await {
        Ok(stream) => {
            //debug!("connected to {}", addr);
            Ok(stream)
        }
        Err(e) => {
            if e.kind() != io::ErrorKind::TimedOut {
                //error!("unknown connect error: '{}'", e);
            }
            Err(e)
        }
    }
}
*/

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let req_str = b"GET / HTTP/1.1
Connection: keep-alive
Host: localhost:9090

";

    // Create the runtime
    //let mut rt = Runtime::new()?;
    let mut rt = runtime::Builder::new().threaded_scheduler()
        .num_threads(15).enable_all().build().unwrap();
    let connection_num = 100;
    let stopped =  Arc::new(AtomicBool::new(false));
    let counter = Arc::new(AtomicI32::new(0));

    // Spawn the root task
    let mut handles = vec![];
    rt.block_on(async {
        (0..connection_num).for_each(|_|{
            let stopped_clone = stopped.clone();
            let counter = counter.clone();
            let mut counter_inner = 0;
            let h = tokio::spawn(async move {
                let mut read_buffer = [0u8; 1024];
                let mut stream = TcpStream::connect("127.0.0.1:9090").await.unwrap();

                while !stopped_clone.load(Ordering::Relaxed) {
                    match stream.write(req_str).await {
                        Ok(_) => match stream.read(&mut read_buffer).await {
                            Ok(n) => {
                                counter_inner += 1;
                                //recv_bytes_total += n;
                                //println!( "read {} bytes, {:?}", n, String::from_utf8(read_buffer[..n].to_vec()) );
                            }
                            Err(_) => {}
                        },
                        Err(e) => println!("{}", e),
                    };
                }
                counter.fetch_add(counter_inner, Ordering::Relaxed);
            });
            handles.push(h);
        });

        tokio::spawn(async move {
            delay_for(Duration::from_secs(5)).await;
            stopped.store(true, Ordering::Relaxed);
        });
    });

    // 不 join 的话，其实内部的 future 们还没有运行完
    rt.block_on(join_all(handles));
    println!("requests : {:?}", counter);

    Ok(())
}

