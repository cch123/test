package main
import "time"

func main() {
	m := make(map[int]int)
		go func(){
			for i:=0;;i++{
				m[i] = i+2
			}
		}()
	
	for i:=0;i<100;i++{
		go func(){
			_ = len(m)
		}()
	}
	time.Sleep(time.Hour)
}
/*
==================
WARNING: DATA RACE
Read at 0x00c000096060 by goroutine 7:
  main.main.func2()
      /Users/xargin/test/go/map/is_len_of_map_thread_safe.go:14 +0x3d

Previous write at 0x00c000096060 by goroutine 6:
  runtime.mapassign_fast64()
      /usr/local/go/src/runtime/map_fast64.go:92 +0x0
  main.main.func1()
      /Users/xargin/test/go/map/is_len_of_map_thread_safe.go:8 +0x50

Goroutine 7 (running) created at:
  main.main()
      /Users/xargin/test/go/map/is_len_of_map_thread_safe.go:13 +0x84

Goroutine 6 (running) created at:
  main.main()
      /Users/xargin/test/go/map/is_len_of_map_thread_safe.go:6 +0x59
==================
*/
