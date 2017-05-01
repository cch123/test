package main

// 抽时间再写个dfs版的

func main() {
	letterCombinations("23")
}

func letterCombinations(digits string) []string {
	var m = map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var candidate = []string{""}
	for i := 0; i < len(digits); i++ {
		// 第i个数字
		valList := m[digits[i]]
		var newCandidate = []string{}
		for j := 0; j < len(candidate); j++ {
			for k := 0; k < len(valList); k++ {
				tmpStr := string(append([]byte(candidate[j]), valList[k]))
				newCandidate = append(newCandidate, tmpStr)
			}
		}
		candidate = newCandidate
	}
	if len(candidate) == 1 {
		return []string{}
	}
	return candidate
}
