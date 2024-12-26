package main

import (
	"fmt"
	"sort"
)

func duplicateBySorting(s string) map[string]int {
	ss := []rune(s)

	res := make(map[string]int)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	i := 0
	for i < len(ss) {
		count := 1
		for i < len(ss)-1 && ss[i] == ss[i+1] {
			count++
			i++
		}
		if count > 1 {
			res[string(ss[i])] = count
		}
		i++
	}
	return res
}

func duplicateByHashing(s string) map[string]int {
	res := make(map[string]int)
	ss := []rune(s)

	for _, r := range ss {
		if _, ok := res[string(r)]; ok {
			res[string(r)]++
		} else {
			res[string(r)] = 1
		}
	}

	for k := range res {
		if res[k] == 1 {
			delete(res, k)
		}
	}

	return res
}

func main() {
	fmt.Println(duplicateBySorting("hello world"))
	fmt.Println(duplicateByHashing("hello world"))
}
