package main

import (
	"fmt"
	"sort"
)

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
//
//	1 <= strs.length <= 200
//	0 <= strs[i].length <= 200
//	strs[i] consists of only lowercase English letters if it is non-empty.
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	for idx := 0; idx < len(strs); idx++ {
		if len(strs[idx]) == 0 {
			return ""
		}
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	commonPrefix := ""
	for idx := 0; idx < len(strs[0]); idx++ {
		matched := 0
		for _, str := range strs[1:] {
			if strs[0][idx] == str[idx] {
				matched++
				continue
			} else {
				return commonPrefix
			}
		}
		if matched == len(strs)-1 {
			commonPrefix += string(strs[0][idx])
		}
	}

	return commonPrefix

}
