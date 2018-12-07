package main

func main() {

	var aux1 = int32(15)
	aux1 = aux1 - ((aux1 >> 1) & 0x55555555)
	aux1 = (aux1 & 0x33333333) + ((aux1 >> 2) & 0x33333333)
	println(((aux1 + (aux1 >> 4)) & 0x0F0F0F0F))
}
