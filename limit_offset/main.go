package main

import "fmt"

func LimitOffsetLoop[T any](items []T, limit, offset int) {
	total := len(items)

	for {
		if offset >= total || limit == 0 {
			break
		}
		if offset+limit > total {
			limit = total - offset
		}

		fmt.Println(items[offset : limit+offset])
		fmt.Println(limit, offset)
		offset += limit
	}
}

func main() {
	inputs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	LimitOffsetLoop(inputs, 1, 0)
}
