package main

func main() {
	var x = func() func() int {
		var idx = 0
		return func() int {
			idx++
			return idx
		}
	}()
	var z = x
	println(z())
	println(z())
	println(z())
	y := x
	println(y())
	println(y())
	println(y())
	println(y())
}
