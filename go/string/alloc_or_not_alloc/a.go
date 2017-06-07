package main

func TestA() {
	var a = []byte{'a', 'b'}
	var b = map[string]bool{}
	b[string(a)] = true
}

func TestB() {
	var a = []byte{'a', 'b'}
	var b = map[string]bool{}
	var c = string(a)
	b[c] = true
}
