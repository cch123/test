https://github.com/golang/go/issues/26703

不打开 -race 的选项的话，有可能 race 的栈显示不全，如两个 goroutine 发生 race，而实际上只显示其中的一个。

但线上显然是没法打开的。

可以在压力测试中开 -race 进行单独的并发测试

