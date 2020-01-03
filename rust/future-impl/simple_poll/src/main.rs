use std::task::Context;
use futures::task::{waker_ref, ArcWake};
use std::sync::Arc;
use futures::FutureExt;
use std::pin::Pin;
use std::future::Future;

fn main() {
    let mut x = Box::pin(async { "fucker" });
    let t = Arc::new(Task{});
    let waker = waker_ref(&t);
    let ctx = &mut Context::from_waker(&waker);

    let x = Pin::as_mut(&mut x);
    let r = x.poll(ctx);
    dbg!(r);

    let mut x = Box::pin(async { "hello" });
    let t = Arc::new(Task{});
    let waker = waker_ref(&t);
    let ctx = &mut Context::from_waker(&waker);

    let r = x.poll_unpin(ctx);
    dbg!(r);
}

// https://www.reddit.com/r/rust/comments/bu0dn4/how_to_poll_other_future_from_the_poll_fn_in/

pub struct Task {}

impl ArcWake for Task {
    fn wake_by_ref(_: &Arc<Self>) {}
}


