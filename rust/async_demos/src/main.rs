use futures::executor::block_on;
use std::pin::Pin;

#[tokio::main]
async fn main() {
    test();
    join_all_async_block().await;
}

fn test() {
    let fut = hello_world();
    // 调用 async 函数有两种方法
    // 1
    // fut.await; -> await 必须在 async block、async function 或者 async closure 中使用
    // 2
    // block_on(fut); -> 如果是非 async 函数，可以直接用 block_on 等待 future 执行完毕
    block_on(fut);
}

async fn hello_world() {
    println!("hello world");
}

async fn join_all_async_block() {
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
