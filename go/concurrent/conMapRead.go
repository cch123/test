package main

//这个例子在我现在这个电脑上试不出崩溃来
//应该是得go 1.6以上才会崩溃吧。
//go run -race xxx.go

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
