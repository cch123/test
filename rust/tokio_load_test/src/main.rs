use tokio::net::TcpStream;
use tokio::prelude::*;
use tokio::signal;
use tokio::time::delay_for;

use std::sync::atomic::AtomicI32;
use std::sync::{
    atomic::{AtomicBool, Ordering},
    Arc,
};

use futures::future::join_all;

use std::time::{Duration, Instant};

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

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let req_str = b"GET / HTTP/1.1
Host: localhost:9090

";
    //Connection: keep-alive

    // Create the runtime
    //let mut rt = Runtime::new()?;
    /*let mut rt = runtime::Builder::new()
    .threaded_scheduler()
    .num_threads(15)
    .enable_all()
    .build()
    .unwrap();
    */


    let connection_num = 100i32;
    let stopped = Arc::new(AtomicBool::new(false));

    let (counter, bytes_counter) = (Arc::new(AtomicI32::new(0)), Arc::new(AtomicI32::new(0)));
    let now = Instant::now();

    // for timeout
    let stopped_c = stopped.clone();
    tokio::spawn(async move {
        delay_for(Duration::from_secs(5)).await;
        stopped_c.store(true, Ordering::SeqCst);
    });

    // for signal
    let stopped_s = stopped.clone();
    tokio::spawn(async move {
        signal::ctrl_c().await.unwrap();
        stopped_s.store(true, Ordering::SeqCst);
    });

    // for requests
    //let (mut tx, mut rx) = mpsc::channel(connection_num);
    let mut handles = vec![];

    (0..connection_num).for_each(|_| {
        //let tx = tx.clone();
        let (counter, bytes_counter) = (counter.clone(), bytes_counter.clone());
        let (mut counter_inner, mut bytes_counter_inner) = (0, 0);
        let stopped_clone = stopped.clone();
        let h = tokio::spawn(async move {
            let mut read_buffer = [0u8; 1024];
            let mut stream = TcpStream::connect("127.0.0.1:9090").await.unwrap();

            while !stopped_clone.load(Ordering::SeqCst) {
                match stream.write(req_str).await {
                    Ok(_) => match stream.read(&mut read_buffer).await {
                        Ok(n) => {
                            counter_inner += 1;
                            bytes_counter_inner += n as i32;
                            //println!( "read {} bytes, {:?}", n, String::from_utf8(read_buffer[..n].to_vec()) );
                        }
                        Err(_) => {}
                    },
                    Err(e) => println!("{}", e),
                };
            }

            // stats update
            counter.fetch_add(counter_inner, Ordering::Relaxed);
            bytes_counter.fetch_add(bytes_counter_inner, Ordering::Relaxed);

        });
        handles.push(h);
    });

    join_all(handles).await;

    println!("{:?}", now.elapsed());

    // 不 join 的话，其实内部的 future 们还没有运行完
    println!("requests : {:?}; bytes : {:?}", counter, bytes_counter);

    Ok(())
}
