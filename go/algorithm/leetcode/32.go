package main

import "fmt"

func main() {
	//)()()))
	//(()
	fmt.Println(longestValidParentheses(`((()))`))
	fmt.Println(longestValidParentheses(`((()))((()))`))

	fmt.Println(longestValidParentheses(`()()`))
	fmt.Println(longestValidParentheses(`()(()`))
}

func longestValidParentheses(s string) int {

	var maxDistance = 0
	var maxLenStack []int
	var parenStack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			parenStack = append(parenStack, '(')
			maxLenStack = append(maxLenStack, 0)
		}

		if s[i] == ')' {
			if len(parenStack) > 0 && parenStack[len(parenStack)-1] == '(' {
				// pop and calc
				parenStack = parenStack[:len(parenStack)-1]
				maxLenStack[len(maxLenStack)-1] += 2
				if maxLenStack[len(maxLenStack)-1] > 2 {
					if len(maxLenStack) >= 2 {
						if maxLenStack[len(maxLenStack)-2] == 0 {
							maxLenStack[len(maxLenStack)-2] = maxLenStack[len(maxLenStack)-1]
							maxLenStack = maxLenStack[0 : len(maxLenStack)-1]
						}
					}
				}

				for len(maxLenStack) >= 2 {
					// merge down to bottom
					if maxLenStack[len(maxLenStack)-2] > 0 {
						maxLenStack[len(maxLenStack)-2] += maxLenStack[len(maxLenStack)-1]
						// pop
						maxLenStack = maxLenStack[:len(maxLenStack)-1]
					} else {
						break
					}
				}
			} else {
				maxLenStack = append(maxLenStack, 0)
				parenStack = append(parenStack, ')')
			}
		}
		if len(maxLenStack) > 0 && maxLenStack[len(maxLenStack)-1] > maxDistance {
			maxDistance = maxLenStack[len(maxLenStack)-1]
		}
		//fmt.Println(parenStack, maxLenStack, maxDistance)
	}

	return maxDistance

}

//0,1,0,1,2,1
