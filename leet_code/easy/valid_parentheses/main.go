package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
//	Open brackets must be closed by the same type of brackets.
//	Open brackets must be closed in the correct order.
//	Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
//
// Output: true
//
// Example 2:
//
// Input: s = "()[]{}"
//
// Output: true
//
// Example 3:
//
// Input: s = "(]"
//
// Output: false
//
// Example 4:
//
// Input: s = "([])"
//
// Output: true
//
// Constraints:
//
//	1 <= s.length <= 104
//	s consists of parentheses only '()[]{}'.
func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("]]"))
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	parenthesisMapper := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			parenthesisMapper = append(parenthesisMapper, ')')
		} else if s[i] == '{' {
			parenthesisMapper = append(parenthesisMapper, '}')
		} else if s[i] == '[' {
			parenthesisMapper = append(parenthesisMapper, ']')
		} else {
			if len(parenthesisMapper) == 0 {
				return false
			}

			firstClosing := parenthesisMapper[len(parenthesisMapper)-1]
			if rune(s[i]) != firstClosing {
				return false
			}
			if len(parenthesisMapper) > 0 {
				parenthesisMapper = parenthesisMapper[:len(parenthesisMapper)-1]
			}
		}
	}
	if len(parenthesisMapper) != 0 {
		return false
	}

	return true
}
