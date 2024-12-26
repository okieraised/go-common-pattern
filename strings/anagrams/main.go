package main

import (
	"fmt"
	"sort"
)

func naiveAreAnagram(s1, s2 string) bool {
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	sort.Slice(ss1, func(i, j int) bool { return ss1[i] < ss1[j] })
	sort.Slice(ss2, func(i, j int) bool { return ss2[i] < ss2[j] })

	return string(ss1) == string(ss2)
}

func hashMapAreAnagram(s1, s2 string) bool {
	charCount := make(map[rune]int)
	for _, char := range []rune(s1) {
		charCount[char]++
	}

	for _, char := range []rune(s2) {
		charCount[char]--
	}

	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}
	return true
}

func freqArrAreAnagram(s1, s2 string) bool {
	freq := make([]int, 26)

	for _, char := range []rune(s1) {
		freq[char-'a']++
	}

	for _, char := range []rune(s2) {
		freq[char-'a']--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "geeks"
	s2 := "kseeg"

	fmt.Println("naive:", naiveAreAnagram(s1, s2))

	fmt.Println("hash:", hashMapAreAnagram(s1, s2))

	fmt.Println("freq:", hashMapAreAnagram(s1, s2))

}
