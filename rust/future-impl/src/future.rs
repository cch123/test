use std::pin::Pin;
use std::task::Waker;
use std::marker::PhantomData;

pub struct Context<'a> {
    waker: &'a Waker,
    _marker: PhantomData<fn(&'a ()) -> &'a ()>,
}

pub trait Future {
    /// The type of value produced on completion.
    type Output;
    fn poll(self: Pin<&mut Self>, cx: &mut Context<'_>) -> Poll<Self::Output>;
}

pub enum Poll<T> {
    Ready(
        T
    ),
    Pending,
}

impl<T> Poll<T> {
    pub fn map<U, F>(self, f: F) -> Poll<U>
        where F: FnOnce(T) -> U
    {
        match self {
            Poll::Ready(t) => Poll::Ready(f(t)),
            Poll::Pending => Poll::Pending,
        }
    }

    /// Returns `true` if this is `Poll::Ready`
    pub fn is_ready(&self) -> bool {
        match *self {
            Poll::Ready(_) => true,
            Poll::Pending => false,
        }
    }

    /// Returns `true` if this is `Poll::Pending`
    pub fn is_pending(&self) -> bool {
        !self.is_ready()
    }
}

impl<T, E> Poll<Result<T, E>> {
    /// Changes the success value of this `Poll` with the closure provided.
    pub fn map_ok<U, F>(self, f: F) -> Poll<Result<U, E>>
        where F: FnOnce(T) -> U
    {
        match self {
            Poll::Ready(Ok(t)) => Poll::Ready(Ok(f(t))),
            Poll::Ready(Err(e)) => Poll::Ready(Err(e)),
            Poll::Pending => Poll::Pending,
        }
    }

    /// Changes the error value of this `Poll` with the closure provided.
    pub fn map_err<U, F>(self, f: F) -> Poll<Result<T, U>>
        where F: FnOnce(E) -> U
    {
        match self {
            Poll::Ready(Ok(t)) => Poll::Ready(Ok(t)),
            Poll::Ready(Err(e)) => Poll::Ready(Err(f(e))),
            Poll::Pending => Poll::Pending,
        }
    }
}

impl<T, E> Poll<Option<Result<T, E>>> {
    /// Changes the success value of this `Poll` with the closure provided.
    pub fn map_ok<U, F>(self, f: F) -> Poll<Option<Result<U, E>>>
        where F: FnOnce(T) -> U
    {
        match self {
            Poll::Ready(Some(Ok(t))) => Poll::Ready(Some(Ok(f(t)))),
            Poll::Ready(Some(Err(e))) => Poll::Ready(Some(Err(e))),
            Poll::Ready(None) => Poll::Ready(None),
            Poll::Pending => Poll::Pending,
        }
    }

    /// Changes the error value of this `Poll` with the closure provided.
    pub fn map_err<U, F>(self, f: F) -> Poll<Option<Result<T, U>>>
        where F: FnOnce(E) -> U
    {
        match self {
            Poll::Ready(Some(Ok(t))) => Poll::Ready(Some(Ok(t))),
            Poll::Ready(Some(Err(e))) => Poll::Ready(Some(Err(f(e)))),
            Poll::Ready(None) => Poll::Ready(None),
            Poll::Pending => Poll::Pending,
        }
    }
}

pub fn block_on<F: Future>(f: F) -> F::Output {
    //pin_mut!(f);
    //run_executor(|cx| f.as_mut().poll(cx))
}


/*
fn run_executor<T, F: FnMut(&mut Context<'_>) -> Poll<T>>(mut f: F) -> T {
    /*
    let _enter = enter().expect(
        "cannot execute `LocalPool` executor from within \
         another executor",
    );

    CURRENT_THREAD_NOTIFY.with(|thread_notify| {
        let waker = waker_ref(thread_notify);
        let mut cx = Context::from_waker(&waker);
        loop {
            if let Poll::Ready(t) = f(&mut cx) {
                return t;
            }
            // Consume the wakeup that occurred while executing `f`, if any.
            let unparked = thread_notify.unparked.swap(false, Ordering::Acquire);
            if !unparked {
                // No wakeup occurred. It may occur now, right before parking,
                // but in that case the token made available by `unpark()`
                // is guaranteed to still be available and `park()` is a no-op.
                thread::park();
                // When the thread is unparked, `unparked` will have been set
                // and needs to be unset before the next call to `f` to avoid
                // a redundant loop iteration.
                thread_notify.unparked.store(false, Ordering::Release);
            }
        }
    })
    */
}
*/