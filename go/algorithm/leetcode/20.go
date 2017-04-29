package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)])"))
}

func isValid(s string) bool {

	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			if s[i] == ']' || s[i] == '}' || s[i] == ')' {
				return false
			}
			stack = append(stack, s[i])

		} else if stack[len(stack)-1] == '{' {
			if s[i] == ']' || s[i] == ')' {
				return false
			}

			if s[i] == '}' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else if stack[len(stack)-1] == '[' {
			if s[i] == '}' || s[i] == ')' {
				return false
			}
			if s[i] == ']' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else if stack[len(stack)-1] == '(' {
			if s[i] == '}' || s[i] == ']' {
				return false
			}
			if s[i] == ')' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
