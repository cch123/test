#![feature(async_closure)]

use futures::executor::block_on;
use futures::{FutureExt, StreamExt};
use std::pin::Pin;
use tokio::sync::mpsc;
use tokio::sync::oneshot;

#[tokio::main]
async fn main() {
    // await、join 执行只能在 async 函数中
    // 如果想在非 async 函数中执行 future，需要用特殊手段
    execute_future_in_non_async_functions();
    join_all_async_block_in_various_ways().await;
    async_block_and_async_fn_are_alike().await;

    futures_can_be_selected().await;

    // producer - consumer
    channel_like_go_buffered().await;
    channel_like_go_oneshot().await;

    basic_synchronization().await;
    advanced_synchronization().await;

    // TODO
    select_examples().await;
}

async fn select_examples() {
    select_all_demo().await;
    select_ok_demo().await;
    select_macro().await;
}

// after either one succ, return
// TODO
async fn select_macro() {
    let (mut tx, mut rx) = mpsc::channel(10);
    let (mut tx1, mut tx2) = (tx.clone(), tx.clone());
    let (mut tx3, mut tx4) = (tx.clone(), tx.clone());
    let (mut producer1, mut producer2, mut producer3, mut producer4) = (
        Box::pin(async {
            // request baidu.com and send result to channel
            // TODO
            tx1.send(111).await.unwrap();
        })
        .fuse(),
        Box::pin(async {
            // request jd.com and send result to channel
            // TODO
            tx2.send(222).await.unwrap();
        })
        .fuse(),
        Box::pin(async {
            // request jd.com and send result to channel
            // TODO
            tx3.send(333).await.unwrap();
        })
        .fuse(),
        Box::pin(async {
            // 在 future 里 sleep 不应该用阻塞操作
            // https://github.com/tokio-rs/tokio/blob/master/tokio/src/time/delay.rs
            tokio::timer::delay_for(std::time::Duration::from_secs(10)).await;
            tx4.send(444).await.unwrap();
        })
        .fuse(),
    );

    // request system 1
    // request system 2
    // select one producer
    futures::select! {
        r = producer1 => {
            println!("producer 1 ready {:?}", r);
        },
        t = producer3 => {
            println!("producer 3 ready {:?}", t);
        },
        s = producer2 => {
            println!("producer 2 ready {:?}", s);
        },
        u = producer4 => {
            println!("producer 4 ready {:?}", u);
        },
    }

    match rx.recv().await {
        Some(msg) => println!("receive {} in consumer", msg),
        None => println!("receive chan closed"),
    }
}

async fn select_all_demo() {
    let a = async { 10 };
    let b = async { 123 };
    let fut = futures::future::select_all(vec![
        Box::pin(a) as Pin<Box<dyn futures::Future<Output = i32>>>,
        Box::pin(b) as Pin<Box<dyn futures::Future<Output = i32>>>,
    ]);

    match fut.await {
        (res, siz, v) => {
            println!("in select all, go res {}", res);
        }
    }
}

async fn select_ok_demo() {
    let a = async { Ok(1) };
    let b = async { Ok(1) };
    let fut = futures::future::select_ok(vec![
        Box::pin(a) as Pin<Box<dyn futures::Future<Output = Result<i32, i32>>>>,
        Box::pin(b) as Pin<Box<dyn futures::Future<Output = Result<i32, i32>>>>,
    ]);

    match fut.await {
        Ok((res, v)) => {
            println!("in select ok, go res {}", res);
        }
        Err(e) => {
            println!("go err in select ok, {:?}", e);
        }
    }
}

async fn try_join_demo() {
    //futures::future::try_join()
}

async fn advanced_synchronization() {
    let mut xx = std::collections::HashMap::new();
    let (tx, mut rx) = mpsc::channel(10);
    // way 2:
    for i in 0..10 {
        let mut tx = tx.clone();
        tokio::spawn(async move {
            // 这里如果不 move 就会告诉你外层的 tx not live long enough
            match tx.send((i, i + 1)).await {
                Ok(_) => {}
                Err(e) => {
                    println!("receiver dropped {}", e);
                    return;
                }
            }
            // Ok 返回一个 ()，一般也不关心，可以用下面这种写法
            /*
            if let Err(_) = tx.send((i, i+1)).await {
            }
            */
        });
    }

    // 这里不 drop 的话，这个 channel 没法退出？
    drop(tx);

    // 然而用 tokio::spawn 的话，似乎就不太容易等待所有 future 都完成后再一并退出了
    // 如果想进行同步，使用 mpsc 是可以的
    // 不用 mpsc 的话
    // 怎么样等待所有 spawn 的 task 结束？
    loop {
        match rx.recv().await {
            Some(k) => {
                println!("got k {:?}", k);
                xx.insert(k.0, k.1);
            }
            None => {
                println!("got none");
                break;
            }
        }
    }
    println!("xx is {:?}", xx);
}

async fn basic_synchronization() {
    // mutex
    let x = std::sync::Arc::new(std::sync::RwLock::new(std::collections::HashMap::new()));
    // spawn 和 block_on 的区别:
    // spawn 类似于 go func
    // 如果不显式同步，那最终不一定会得到执行

    let mut futs = vec![];
    for i in 0..20 {
        let y = std::sync::Arc::clone(&x);
        // 这里看起来有两种写法
        // way 1:
        futs.push(async move {
            y.write().unwrap().insert(i as i32, 3);
            println!("{:?}", y.read().unwrap());
        });
    }

    // wg.Wait()
    block_on(futures::future::join_all(futs));

    // spawn 不一定都能得到执行，所以这里理论上应该每次结果都是一样的
    println!("x is : {:?}", x.read().unwrap());
}

async fn channel_like_go_oneshot() {
    let (tx, rx) = oneshot::channel();

    tokio::spawn(async move {
        if let Err(_) = tx.send("hell") {
            println!("the receiver dropped");
        }
    });

    match rx.await {
        Ok(v) => println!("got = {:?}", v),
        Err(_) => println!("the sender dropped"),
    }
}

async fn channel_like_go_buffered() {
    let (mut tx, mut rx) = mpsc::channel(100);

    tokio::spawn(async move {
        for i in 0..10 {
            if let Err(_) = tx.send(i).await {
                println!("receiver dropped");
                return;
            }
        }
    });

    loop {
        // 当所有 sender 都被 drop 之后，从 channel 中就会收到 None
        // 不像 go 那样需要主动关闭 channel 或者 for range
        // Receive the next value for this receiver.
        // `None` is returned when all `Sender` halves have dropped, indicating
        // that no further values can be sent on the channel.
        match rx.recv().await {
            Some(i) => {
                // like go func here
                tokio::spawn(async move {
                    println!("got = {} in spawned coroutine", i);
                });
            }
            None => {
                println!("got none");
                break;
            }
        }
    }
}

// select 在其中一个 future ready 后即返回
// 可以用来优化长尾请求
async fn futures_can_be_selected() {
    let a = async { 1 };
    let b = async { 3 };
    let c = async { 5 };
    let (mut x, mut y, mut z) = (Box::pin(a).fuse(), Box::pin(b).fuse(), Box::pin(c).fuse());
    loop {
        futures::select! {
            xx = x => {
                println!("a ready {}", xx);
            },
            yy = y => {
                println!("b ready {}", yy);
            },
            zz = z => {
                println!("c ready {}", zz);
            },
            complete => {
                println!("a,b,c all complete");
                break;
            },
            default => unreachable!(),
        }
    }

    /*
    这样是不行的，会认为在 loop 过程中，a，b、c 发生了 move
    loop {
        futures::select! {
            xx = Box::pin(a).fuse() => {
                println!("a ready {}", xx);
            },
            yy = Box::pin(b).fuse() => {
                println!("b ready {}", yy);
            },
            zz = Box::pin(c).fuse() => {
                println!("c ready {}", zz);
            },
            complete => {
                println!("a,b,c all complete");
                break;
            },
            default => unreachable!(),
        }
    }

    */

    // https://rust-lang.github.io/async-book/06_multiple_futures/03_select.html
}

fn execute_future_in_non_async_functions() {
    let fut = hello_world();
    // 调用 async 函数有两种方法
    // 1
    // fut.await; -> await 必须在 async block、async function 或者 async closure 中使用
    // 2
    // block_on(fut); -> 如果是非 async 函数，可以直接用 block_on 等待 future 执行完毕
    // block_on 运行 future 一样可以获取到 future 的 Output
    let res = block_on(fut);
    dbg!(res);

    // 自己主动 poll 就麻烦多了，实际上还是得实现一个类似上面 block_on 功能的函数
    // https://users.rust-lang.org/t/how-to-wait-an-async-in-non-async-function/28388/21
    // 尝试让下面这段代码通过编译
    /*
    let poll_result = fut.poll();
    loop {
        match poll_result {
            std::task::Poll::Ready(v) => {
                println!("yes poll ready {}", v);
                break;
            },
            std::task::Poll::Pending => {
                println!("not ready yet");
            }
        }
    }
    */
}

async fn return_string() -> String {
    "def".to_string()
}

async fn async_block_and_async_fn_are_alike() {
    let a = async { "abc".to_string() }; // this async block returns impl Future<Output=String>
    let b = return_string(); // this async func also returns impl Future<Output=String>

    // async closure 现在还是 unstable，需要开 feature gate
    // #![feature(async_closure)]
    let f = async || -> String { "ggg".to_string() };
    let c = f();

    dbg!(futures::future::join3(a, b, c).await);
}

async fn hello_world() -> i32 {
    println!("hello world");
    250
}

async fn join_all_async_block_in_various_ways() {
    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };

    // join 宏相当于 a.await，b.await，c.await，d.await
    match futures::join!(a, b, c, d) {
        (1, 2, 3, 4) => println!("as expected"),
        (_, _, _, _) => println!("oh no"),
    }

    // future 库中还提供了一个 join_all 的函数，但 join_all 使用较麻烦
    // 例如像下面这样的不行的
    // futures::future::join_all(vec![a, b, c, d]).await;
    // ----> 会报：
    /*
        error[E0308]: mismatched types
      --> src/main.rs:35:39
       |
    35 |     futures::future::join_all(vec![a, b, c, d]).await;
       |                                       ^ expected generator, found a different generator
       |
       = note: expected type `impl core::future::future::Future` (generator)
                  found type `impl core::future::future::Future` (generator)
        不限制多个 future 类型的可以用 join，join2，join3，join4，join5
        很丑
        */
    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };
    dbg!(futures::future::join4(a, b, c, d).await);

    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };
    let res = futures::future::join_all(vec![
        Box::pin(a) as Pin<Box<dyn futures::Future<Output = i32>>>,
        Box::pin(b) as Pin<Box<dyn futures::Future<Output = i32>>>,
        Box::pin(c) as Pin<Box<dyn futures::Future<Output = i32>>>,
        Box::pin(d) as Pin<Box<dyn futures::Future<Output = i32>>>,
    ])
    .await;
    dbg!(res);
}
