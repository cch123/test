package main

//这个例子需要在go 1.6以上的版本才会直接崩溃
//可以使用
//go run -race xxx.go
//来检测程序的数据竞争

//import "time"

func main() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"] = 1
			//		time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			//		time.Sleep(time.Microsecond)
		}
	}()
	select {}
}
