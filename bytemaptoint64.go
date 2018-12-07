package main

func main() {
	var b = []byte{2, 1}
	println(bytes2int64(b))
}

func bytes2int64(bList []byte) int64 {
	var res int64
	for i := 0; i < len(bList); i++ {
		res = res*128 + int64(bList[i])
	}
	return res
}
