package main

import "fmt"
import "strings"

func main() {
	a := `Let's take LeetCode contest`
	fmt.Println(reverseWords(a))
}

func reverseWords(s string) string {
	strArr := strings.Split(s, " ")
	var reversedStrArr []string
	for _, v := range strArr {
		byteArr := []byte(v)
		for i := 0; i < len(v)/2; i++ {
			byteArr[i], byteArr[len(v)-1-i] = byteArr[len(v)-1-i], byteArr[i]
		}
		reversedStrArr = append(reversedStrArr, string(byteArr))
	}
	return strings.Join(reversedStrArr, " ")
}
