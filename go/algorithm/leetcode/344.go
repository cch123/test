package main

func main() {
}

func reverseString(s string) string {
	var res []byte

	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	return string(res)
}
