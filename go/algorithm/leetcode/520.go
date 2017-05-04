package main

func main() {
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	var big = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var little = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	var bigMap = make(map[byte]struct{})
	var littleMap = make(map[byte]struct{})
	for i := 0; i < len(big); i++ {
		bigMap[big[i]] = struct{}{}
		littleMap[little[i]] = struct{}{}
	}

	if _, ok := littleMap[word[1]]; ok {
		// little mode, from 1 to last is little
		for i := 2; i < len(word); i++ {
			if _, ok := bigMap[word[i]]; ok {
				return false
			}
		}
	} else {
		// big mode, all big is valid
		for i := 0; i < len(word); i++ {
			if _, ok := littleMap[word[i]]; ok {
				return false
			}
		}
	}
	return true
}
