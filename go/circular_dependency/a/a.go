package a

import "../b"

func Test() {
	b.FormatB()
	println("test a")
}

func FormatA() {
	println("format a")
}
