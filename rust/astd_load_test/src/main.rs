use async_std::io;
use async_std::net::TcpStream;
use async_std::prelude::*;
use async_std::sync::channel;
use async_std::task;
use futures::executor::LocalPool;
use futures::task::SpawnExt;
use futures_timer::Delay;
use std::net::{SocketAddr, ToSocketAddrs};
use std::time::Duration;
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;

async fn connect(addr: &SocketAddr) -> io::Result<TcpStream> {
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

fn main() {
    let thread_num = 10;
    let connection_per_thread = 10;
    let addr = "localhost:9090".to_socket_addrs().unwrap().next().unwrap();
    let req_str = b"GET / HTTP/1.1
Connection: keep-alive
Host: localhost:9090

";
    let (sender, receiver) = channel(thread_num * connection_per_thread);

    let mut thread_handle_list = vec![];

    (0..thread_num).for_each(|_| {
        let sender = sender.clone();
        let h = std::thread::spawn(move || {
            let mut pool = LocalPool::new();
            let spawner = pool.spawner();
            (0..connection_per_thread).for_each(|_| {
                let sender = sender.clone();
                let (mut recv_bytes_total, mut counter) = (0, 0);
                let stopped =  Arc::new(AtomicBool::new(false));

                // task
                let stopped_clone = stopped.clone();
                spawner
                    .spawn(async move {
                        let mut read_buffer = [0u8; 1024];
                        let mut stream = connect(&addr).await.unwrap();

                        while !stopped_clone.load(Ordering::Relaxed) {
                            counter += 1;
                            match stream.write_all(req_str).await {
                                Ok(_) => match stream.read(&mut read_buffer).await {
                                    Ok(n) => {
                                        recv_bytes_total += n;
                                        //println!( "read {} bytes, {:?}", n, String::from_utf8(read_buffer[..n].to_vec()) );
                                    }
                                    Err(_) => {}
                                },
                                Err(e) => println!("{}", e),
                            };
                        }
                        println!("{},{}", counter, recv_bytes_total);
                        sender.send((counter, recv_bytes_total)).await;
                    })
                    .unwrap();

                // timer
                let stopped_clone = stopped.clone();
                spawner
                    .spawn(async move {
                        Delay::new(Duration::from_secs(5)).await;
                        stopped_clone.store(true, Ordering::Relaxed);
                    })
                    .unwrap();
            });
            drop(sender);
            pool.run();
        });
        thread_handle_list.push(h);
    });
    drop(sender);

    for h in thread_handle_list {
        h.join().unwrap();
    }

    let mut total_bytes = 0;
    let mut total_requests = 0;
    task::block_on(async {
        loop {
            match receiver.recv().await {
                Some(elem) => {
                    println!("receiver {:?}", elem);
                    total_bytes += elem.1;
                    total_requests += elem.0;
                }
                None => {
                    println!("channel closed");
                    break;
                }
            }
        }
    });
    println!(
        "total bytes : {};\ntotal requests : {};",
        total_bytes, total_requests
    );
    /*
      // 看一下为什么这里会报：
      error[E0507]: cannot move out of a shared reference
    --> src/main.rs:32:9
      thread_handle_list.iter().for_each(|h|{
          //h.join().unwrap();
          h.join().unwrap();
      });
      */
      report()
}

// https://docs.rs/signal-hook/0.1.12/signal_hook/
fn report() {
    /*
    Running 5s test @ http://localhost:9090
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.38ms   64.61ms 413.27ms   89.43%
    Req/Sec     4.81k     2.12k   11.52k    78.18%
  26462 requests in 606.64ms, 2.83MB read
Requests/sec:  43620.60
Transfer/sec:      4.66MB
    */
}
