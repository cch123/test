use futures::future::BoxFuture;
use std::sync::mpsc::{sync_channel, Receiver, SyncSender};
use std::sync::{Arc, Mutex};
use std::future::Future;
use futures::task::{waker_ref, ArcWake};
use std::task::{Context, Poll};


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
        self.task_sender.send(task).expect("too many tasks queued");
    }
}

pub struct Task {
    future: Mutex<Option<BoxFuture<'static, ()>>>,
    task_sender: SyncSender<Arc<Task>>,
}

impl ArcWake for Task {
    fn wake_by_ref(arc_self: &Arc<Self>) {
        let cloned = arc_self.clone();
        arc_self.task_sender.send(cloned).expect("too many tasks queued");
    }
}

pub struct Executor {
    ready_queue: Receiver<Arc<Task>>,
}

impl Executor {
    pub fn run(&self) {
        while let Ok(task) = self.ready_queue.recv() {
            let mut future_slot = task.future.lock().unwrap();
            if let Some(mut fut) = future_slot.take() {
                let waker = waker_ref(&task);
                let context = &mut Context::from_waker(&*waker);

                match fut.as_mut().poll(context) {
                    Poll::Pending => {
                        *future_slot = Some(fut);
                    }
                    _ => {} // do nothing?
                }
            }
        }
    }
}

pub fn new_executor_and_spawner() -> (Executor, Spawner) {
    const MAX_QUEUE_TASKS: usize = 10_000;
    let (task_sender, ready_queue) = sync_channel(MAX_QUEUE_TASKS);
    (Executor { ready_queue }, Spawner { task_sender })
}
