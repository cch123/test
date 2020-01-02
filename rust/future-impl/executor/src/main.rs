pub mod executor;

fn main() {
    let (executor, spawner) = executor::new_executor_and_spawner();

    spawner.spawn(async {
        println!("hello in async 1");
        println!("world in async 2");
    });

    drop(spawner);
    executor.run();
}
