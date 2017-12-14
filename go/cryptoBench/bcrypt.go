package main

import "fmt"
import "golang.org/x/crypto/bcrypt"
import "time"
import "runtime"

func test(cost int, id int) {
	fmt.Println(time.Now(), "BEGIN test ", id)
	code, _ := bcrypt.GenerateFromPassword([]byte("password"), cost)
	fmt.Println(time.Now(), "END test ", id, code[0])
}

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		//fmt.Println(time.Now(), "BEGIN test ", i)
		//code, _ := bcrypt.GenerateFromPassword([]byte("password"), 17)
		//fmt.Println(time.Now(), "END test ", i, code[0])
		go test(15, i)
	}
	time.Sleep(1e16)
}
