use futures::{
    future::{BoxFuture, FutureExt},
    task::{waker_ref, ArcWake},
};
use std::future::Future;

use std::sync::{Arc, Mutex};
use std::time::Duration;

pub mod executor;

fn main() {
    println!("Hello, world!");
}
