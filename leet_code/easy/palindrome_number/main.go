package main

import (
	"fmt"
)

// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.
//
// Example 1:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
//
// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//
// Example 3:
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
// Constraints:
//
//	-231 <= x <= 231 - 1
//
// Follow up: Could you solve it without converting the integer to a string?
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(99))
	fmt.Println(isPalindrome(4294967295))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var digits []int
	for x > 0 {
		nthDigit := x % 10
		x = x / 10
		digits = append(digits, nthDigit)
	}

	start := 0
	end := len(digits) - 1

	for start < end {
		if digits[start] != digits[end] {
			return false
		}
		start++
		end--
	}

	return true
}
