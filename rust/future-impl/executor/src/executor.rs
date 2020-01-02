use futures::future::BoxFuture;
use std::sync::mpsc::{sync_channel, Receiver, SyncSender};
use std::sync::{Arc, Mutex};
use std::future::Future;

pub struct Spawner {
    task_sender: SyncSender<Arc<Task>>,
}

impl Spawner {
    pub fn spawn(&self, future: impl Future<Output = ()> + 'static + Send) {
        let fut = Box::pin(future);
        let task = Arc::new(Task {
            future: Mutex::new(Some(fut)),
            task_sender: self.task_sender.clone(),
        });
    }
}

pub struct Task {
    future: Mutex<Option<BoxFuture<'static, ()>>>,
    task_sender: SyncSender<Arc<Task>>,
}

pub struct Executor {
    ready_queue: Receiver<Arc<Task>>,
}

pub fn new_executor_and_spawner() -> (Executor, Spawner) {
    const MAX_QUEUE_TASKS: usize = 10_000;
    let (task_sender, ready_queue) = sync_channel(MAX_QUEUE_TASKS);
    (Executor { ready_queue }, Spawner { task_sender })
}
