只要简单的import _ "net/http/pprof"

然后就可以localhost:port/debug/pprof来看了


// cpu gvz
go tool pprof ../producer_consumer ./cpu.prof
