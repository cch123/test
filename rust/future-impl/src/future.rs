/*
Futures can be advanced by calling the poll function,
which will drive the future as far towards completion as possible.
If the future completes,
it returns Poll::Ready(result).
If the future is not able to complete yet,
it returns Poll::Pending and arranges for the wake() function
to be called when the Future is ready to make more progress.
When wake() is called,
the executor driving the Future will call poll again
so that the Future can make more progress.
 */
trait SimpleFuture {
    type Output;
    /*
    Without wake(),
    the executor would have no way of knowing
    when a particular future could make progress,
    and would have to be constantly polling every future.
    With wake(),
    the executor knows exactly which futures are ready to be polled.
    */
    fn poll(&mut self, wake: fn()) -> Poll<Self::Output>;
}

enum Poll<T> {
    Ready(T),
    Pending,
}

struct Socket {}
impl Socket {
    pub fn has_data_to_read(&self) -> bool {
        true
    }

    pub fn read_buf(&self) -> Vec<u8> {
        vec![]
    }

    pub fn set_readable_callback(&self, wake: fn()) {
        println!("when ready call wake");
    }
}

pub struct SocketRead<'a> {
    socket: &'a Socket,
}

impl SimpleFuture for SocketRead<'_> {
    type Output = Vec<u8>;

    fn poll(&mut self, wake: fn()) -> Poll<Self::Output> {
        if self.socket.has_data_to_read() {
            // The socket has data-- read it into a buffer and return it.
            Poll::Ready(self.socket.read_buf())
        } else {
            // The socket does not yet have data.
            //
            // Arrange for `wake` to be called once data is available.
            // When data becomes available, `wake` will be called, and the
            // user of this `Future` will know to call `poll` again and
            // receive data.
            self.socket.set_readable_callback(wake);
            Poll::Pending
        }
    }
}

/// A SimpleFuture that runs two other futures to completion concurrently.
///
/// Concurrency is achieved via the fact that calls to `poll` each future
/// may be interleaved, allowing each future to advance itself at its own pace.
pub struct Join<FutureA, FutureB> {
    // Each field may contain a future that should be run to completion.
    // If the future has already completed, the field is set to `None`.
    // This prevents us from polling a future after it has completed, which
    // would violate the contract of the `Future` trait.
    a: Option<FutureA>,
    b: Option<FutureB>,
}

impl<FutureA, FutureB> SimpleFuture for Join<FutureA, FutureB>
where
    FutureA: SimpleFuture<Output = ()>,
    FutureB: SimpleFuture<Output = ()>,
{
    type Output = ();
    fn poll(&mut self, wake: fn()) -> Poll<Self::Output> {
        // Attempt to complete future `a`.
        if let Some(a) = &mut self.a {
            if let Poll::Ready(()) = a.poll(wake) {
                self.a.take();
            }
        }

        // Attempt to complete future `b`.
        if let Some(b) = &mut self.b {
            if let Poll::Ready(()) = b.poll(wake) {
                self.b.take();
            }
        }

        if self.a.is_none() && self.b.is_none() {
            // Both futures have completed-- we can return successfully
            Poll::Ready(())
        } else {
            // One or both futures returned `Poll::Pending` and still have
            // work to do. They will call `wake()` when progress can be made.
            Poll::Pending
        }
    }
}

/// A SimpleFuture that runs two futures to completion, one after another.
//
// Note: for the purposes of this simple example, `AndThenFut` assumes both
// the first and second futures are available at creation-time. The real
// `AndThen` combinator allows creating the second future based on the output
// of the first future, like `get_breakfast.and_then(|food| eat(food))`.
pub struct AndThenFut<FutureA, FutureB> {
    first: Option<FutureA>,
    second: FutureB,
}

impl<FutureA, FutureB> SimpleFuture for AndThenFut<FutureA, FutureB>
where
    FutureA: SimpleFuture<Output = ()>,
    FutureB: SimpleFuture<Output = ()>,
{
    type Output = ();
    fn poll(&mut self, wake: fn()) -> Poll<Self::Output> {
        if let Some(first) = &mut self.first {
            match first.poll(wake) {
                // We've completed the first future-- remove it and start on
                // the second!
                Poll::Ready(()) => self.first.take(),
                // We couldn't yet complete the first future.
                Poll::Pending => return Poll::Pending,
            };
        }
        // Now that the first future is done, attempt to complete the second.
        self.second.poll(wake)
    }
}

// https://rust-lang.github.io/async-book/02_execution/02_future.html
/*
trait Future {
    type Output;
    fn poll(
        // Note the change from `&mut self` to `Pin<&mut Self>`:
        self: Pin<&mut Self>,
        // and the change from `wake: fn()` to `cx: &mut Context<'_>`:
        cx: &mut Context<'_>,
    ) -> Poll<Self::Output>;
}
The first change you'll notice is that our self type is no longer &mut self,
but has changed to Pin<&mut Self>.
We'll talk more about pinning in a later section,
but for now know that it allows us to create futures that are immovable.
Immovable objects can store pointers between their fields,
e.g. struct MyFut { a: i32, ptr_to_a: *const i32 }.
Pinning is necessary to enable async/await.

Secondly, wake: fn() has changed to &mut Context<'_>.
In SimpleFuture,
we used a call to a function pointer (fn()) to tell the future executor that the future in question should be polled.
However, since fn() is just a function pointer,
it can't store any data about which Future called wake.

In a real-world scenario, a complex application like a web server may have thousands of different connections whose wakeups should all be managed separately. The Context type solves this by providing access to a value of type Waker, which can be used to wake up a specific task.
 */