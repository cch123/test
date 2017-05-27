// in nginx, it append a line feed to each log line
//http://www.perlmonks.org/?node_id=264431
package main

func main() {
	var a = "\x0a"
	var b = "\n"
	println(a == b)
	println([]byte(b)[0])
}
