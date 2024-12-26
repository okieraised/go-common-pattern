package main

import "fmt"

func isPalindromeByTwoPtr(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {

	fmt.Println(isPalindromeByTwoPtr("aa"))
	fmt.Println(isPalindromeByTwoPtr("madam"))
	fmt.Println(isPalindromeByTwoPtr("hello"))
}
