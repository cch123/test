//!
use async_std::{fs::File, prelude::*, sync::channel, task};
use futures::future::{join_all, select_all, select_ok};
use std::pin::Pin;

fn main() {
    read_file();
    block_wait();
    use_async_std_join();
    spawn_blocking();
    task::block_on(async_sleep_and_delay());
    task::block_on(join_mul());
    task::block_on(select_all_demo());
    task::block_on(select_ok_demo());
    task::block_on(spawn_tasks_in_vec());
    use_async_std_select();

    try_join();
    try_race();

    // producer consumer
    task::block_on(mpsc_demo());
    // mpmc
    // TODO
}

async fn mpsc_demo() {
    let (sender, receiver) = channel(3);
    for i in 0..10 {
        let sender = sender.clone();
        task::spawn(async move {
            sender.send(i).await;
        });
    }
    task::sleep(std::time::Duration::from_millis(100)).await;
    // 确保最终 receiver 会退出
    drop(sender);

    loop {
        match receiver.recv().await {
            Some(elem) => println!("receiver {}", elem),
            None => {
                println!("channel closed");
                break;
            },
        }
    }
}

// try race 的 future 必须返回 Result 类型
// 只要其中有一个 future 能成功，就会返回成功的结果
// 全部返回 error 的话，会返回其中的一个 error
fn try_race() {
    let a = async { Ok::<i32, i32>(111) };
    let b = async { Ok::<i32, i32>(333) };
    let c = a.try_race(b);
    let r = task::block_on(c);
    dbg!(r.unwrap());

    let a = async { Ok::<i32, i32>(111) };
    let b = async { Err::<i32, i32>(333) };
    let c = a.try_race(b);
    let r = task::block_on(c);
    dbg!(r);

    let a = async { Err::<i32, i32>(111) };
    let b = async { Err::<i32, i32>(333) };
    let c = a.try_race(b);
    let r = task::block_on(c);
    dbg!(r);
}

// try join 在 future 返回 error 时，会返回一个 error
// future 都成功时，会返回 tuple
// future 全部返回 err 时，只返回第一个 err
fn try_join() {
    let a = async { Ok::<i32, i32>(111) };
    let b = async { Err::<i32, i32>(333) };
    let c = a.try_join(b);
    let r = task::block_on(c);
    match r {
        Ok((x, y)) => println!("{}, {}", x, y),
        Err(e) => println!("err is {}", e),
    }

    let a = async { Ok::<i32, i32>(111) };
    let b = async { Ok::<i32, i32>(333) };
    let c = a.try_join(b);
    let r = task::block_on(c);
    match r {
        Ok((x, y)) => println!("{}, {}", x, y),
        Err(e) => println!("err is {}", e),
    }

    let a = async { Err::<i32, i32>(111) };
    let b = async { Err::<i32, i32>(333) };
    let c = a.try_join(b);
    let r = task::block_on(c);
    match r {
        Ok((x, y)) => println!("{}, {}", x, y),
        Err(e) => println!("err is {}", e),
    }
}

async fn async_sleep_and_delay() {
    task::sleep(std::time::Duration::from_secs(1)).await;
    let f = async { 1 }.delay(std::time::Duration::from_secs(1));
    f.await;
}

// https://docs.rs/async-std/1.0.1/async_std/task/fn.spawn_blocking.html
// Spawns a blocking task.
// The task will be spawned onto a thread pool specifically dedicated to blocking tasks. This is useful to prevent long-running synchronous operations from blocking the main futures executor.
fn spawn_blocking() {
    let h = task::spawn_blocking(|| {
        (0..1000000000).for_each(|_| {});

        1
    });

    dbg!(h);

    //let res = block_on(h);
    //dbg!(res);
}

async fn spawn_tasks_in_vec() {
    let mut v = vec![];
    for i in 0..10 {
        v.push(task::spawn(async move { i }));
    }

    let r = join_all(v).await;
    dbg!(r);
}

// async std 里的 race 和 select 意思差不多
// https://docs.rs/async-std/1.0.1/async_std/future/trait.Future.html#method.race
fn use_async_std_select() {
    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };
    let e = a.race(b);
    let f = c.race(d);
    let r = task::block_on(e.join(f));
    dbg!(r);
}

// join in async std future
fn use_async_std_join() {
    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };

    // 这个破 join 只能 join 两个 Output 相同的 future
    // 垃圾
    let e = a.join(b);
    let f = c.join(d);
    let r = task::block_on(e.join(f));

    dbg!(r);

    let a = async { 1 };
    let b = async { 2 };
    let c = async { 3 };
    let d = async { 4 };
    let r = task::block_on(join_all::<
        Vec<Pin<Box<dyn Future<Output = i32>>>>,
    >(vec![
        Box::pin(a),
        Box::pin(b),
        Box::pin(c),
        Box::pin(d),
    ]));
    dbg!(r);
}

async fn select_ok_demo() {
    let a = async { Ok(1) };
    let b = async { Ok(1) };
    let fut = select_ok::<Vec<Pin<Box<dyn Future<Output = Result<i32, i32>>>>>>(vec![
        Box::pin(a),
        Box::pin(b), // as Pin<Box<dyn futures::Future<Output = Result<i32, i32>>>>,
    ]);

    match fut.await {
        Ok((res, _v)) => {
            println!("in select ok, go res {}", res);
        }
        Err(e) => {
            println!("go err in select ok, {:?}", e);
        }
    }

    let a = task::spawn(async { Ok::<i32, i32>(1) });
    let b = task::spawn(async { Ok::<i32, i32>(1) });
    let fut = select_ok(vec![a, b]);
    match fut.await {
        Ok((res, _v)) => {
            println!("in select ok, go res {}", res);
        }
        Err(e) => {
            println!("go err in select ok, {:?}", e);
        }
    }
}

async fn select_all_demo() {
    let a = async { 10 };
    let b = async { 123 };
    let fut = select_all::<Vec<Pin<Box<dyn Future<Output = i32>>>>>(vec![Box::pin(a), Box::pin(b)]);

    match fut.await {
        (res, _siz, _v) => {
            println!("in select all, go res {}", res);
        }
    }

    let a = task::spawn(async { 10 });
    let b = task::spawn(async { 123 });
    let fut = select_all(vec![a, b]);

    match fut.await {
        (res, _siz, _v) => {
            println!("in select all, go res {}", res);
        }
    }
}

async fn join_mul() {
    let a = async { 1 };
    let b = async { 1 };
    let c = async { 1 };
    let d = async { 1 };
    let res = join_all::<Vec<Pin<Box<dyn Future<Output = i32>>>>>(vec![
        Box::pin(a),
        Box::pin(b),
        Box::pin(c),
        Box::pin(d),
    ])
    .await;
    dbg!(res);

    let a = task::spawn(async { 1 });
    let b = task::spawn(async { 1 });
    let c = task::spawn(async { 1 });
    let d = task::spawn(async { 1 });
    let res = join_all(vec![a, b, c, d]).await;
    dbg!(res);
}

fn block_wait() {
    let a = async { 1 };
    let r = task::block_on(a);
    dbg!(r);
}

fn read_file() {
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
