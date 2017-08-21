package main

import "math/rand"
import "time"

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	for i := 0; i < 100; i++ {
		// 伪随机
		x := rand.Intn(100)
		println("fake", x)

		// 真随机
		y := random.Intn(100)
		println("true", y)
	}
}
