package main

func main() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
