#![feature(async_closure)]

use futures::executor::block_on;
use std::pin::Pin;

#[tokio::main]
async fn main() {
    // await、join 执行只能在 async 函数中
    // 如果想在非 async 函数中执行 future，需要用特殊手段
    execute_future_in_non_async_functions();
    join_all_async_block_in_various_ways().await;
    async_block_and_async_fn_are_alike().await;
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
    let f = async || -> String {
        "ggg".to_string()
    };
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
