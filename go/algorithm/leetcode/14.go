package main

import "fmt"

func main() {

	fmt.Println(longestCommonPrefix([]string{"abc", "abcdefg", "abcd"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	maxPrefix := []byte(strs[0])
	for i := 1; i < len(strs); i++ {
		strLeft := strs[i-1]
		strRight := strs[i]

		currentCommonPrefix := []byte{}
		for j := 0; j < len(strLeft) && j < len(strRight); j++ {
			if strLeft[j] != strRight[j] {
				break
			}

			currentCommonPrefix = append(currentCommonPrefix, strLeft[j])
		}
		if len(currentCommonPrefix) < len(maxPrefix) {
			maxPrefix = currentCommonPrefix
		}
	}
	return string(maxPrefix)

}
