#![feature(async_closure)]
use tokio::sync::mpsc;

#[tokio::main]
async fn main() {
    let (tx, mut rx) = mpsc::channel(10);
    for i in 0..10 {
        let mut tx = tx.clone();
        tokio::spawn(async move {
            match tx.send(i).await {
                Ok(_) => {
                    println!("send times {}", i);
                },
                Err(e) => {
                    println!("send error {}", e);
                },
            }
        });
    }
    drop(tx);

    tokio::timer::delay_for(std::time::Duration::from_secs(1));
    // 不要在 async 函数中使用这种 block 函数
    // std::thread::sleep(std::time::Duration::from_secs(1));

    loop {
        match rx.recv().await {
            Some(msg) => println!("received msg: {}", msg),
            None => {
                println!("done");
                break;
            },
        }
    }
}
