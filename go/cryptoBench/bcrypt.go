package main

import "fmt"
import "golang.org/x/crypto/bcrypt"
import "time"
import "runtime"

func test(task chan struct{}, cost int, id int) {
	for _ = range task {
		runtime.LockOSThread()
		fmt.Println(time.Now(), "BEGIN test ", id)
		code, _ := bcrypt.GenerateFromPassword([]byte("password"), cost)
		fmt.Println(time.Now(), "END test ", id, code[0])
	}
}

func main() {
	task := make(chan struct{}, 4)
	var max = 40
	//fmt.Println(time.Now(), "BEGIN test ", i)
	//code, _ := bcrypt.GenerateFromPassword([]byte("password"), 15)
	//fmt.Println(time.Now(), "END test ", i, code[0])
	go func() {
		for i := 0; i < max; i++ {
			select {
			case task <- struct{}{}:
			default:
			}
		}
	}()

	for i := 0; i < 2; i++ {
		go test(task, 15, i)
	}
	time.Sleep(1e16)
}
