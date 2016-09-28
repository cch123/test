package main

func main() {
	var a = "\346\265\213\350\257\225"
	var b = []rune(a)
	println(a)
	println(len(a), len(b))
	for i := 0; i < len(b); i++ {
		println(string(b[i]))
	}

	var c = []rune("呵呵")
	for i := 0; i < len(c); i++ {
		println(c[i])
	}

	var d = "\u77e5"
	println(d)
}
