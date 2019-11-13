//!
use async_std::{fs::File, prelude::*, task};
use futures::future::join_all;
use futures::future::select_all;
use futures::future::select_ok;
use std::pin::Pin;

fn main() {
    read_file();
    block_wait();
    use_async_std_join();
    task::block_on(join_mul());
    task::block_on(join_mul2());
    task::block_on(select_all_demo());
    task::block_on(select_all_demo2());
    task::block_on(select_ok_demo());
    task::block_on(select_ok_demo2());
    task::block_on(spawn_tasks_in_vec());
    // TODO, try join, try join all, try select
    // TODO, join, select macro
}

async fn spawn_tasks_in_vec() {
    let mut v = vec![];
    for i in 0..10 {
        v.push(task::spawn(async move {
            i
        }));
    }

    let r = futures::future::join_all(v).await;
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
    let r = task::block_on(futures::future::join_all::<
        Vec<Pin<Box<dyn futures::Future<Output = i32>>>>,
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
}

async fn select_ok_demo2() {
    let a = task::spawn(async { Ok::<i32,()>(1) });
    let b = task::spawn(async { Ok::<i32,()>(1) });
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
}

async fn select_all_demo2() {
    let a = task::spawn(async { 10 });
    let b = task::spawn(async { 123 });
    let fut = select_all(vec![a, b]);

    match fut.await {
        (res, _siz, _v) => {
            println!("in select all, go res {}", res);
        }
    }
}

async fn join_mul2() {
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
}

async fn join_mul() {
    let a = task::spawn(async { 1 });
    let b = task::spawn(async { 1 });
    let c = task::spawn(async { 1 });
    let d = task::spawn(async { 1 });
    let e = task::spawn(async { 1 });
    // async std 会直接把 join handle 给返回回来，不像 join_mul2 那样得自己做 Box::pin
    dbg!(e);
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
