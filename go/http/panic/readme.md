虽然 http 库里有对 handlerFunc 的 panic 进行处理，

但实际开发中，如果在 handler 中需要再启动 goroutine，那么新启动的 goroutine 中 panic 还是会导致程序挂掉的。
