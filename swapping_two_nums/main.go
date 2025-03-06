package main

import "fmt"

func main() {
	a := 5
	b := 7

	fmt.Println("1. a:", a, "b:", b)

	a ^= b
	fmt.Println("2. a:", a, "b:", b)
	b ^= a
	fmt.Println("3. a:", a, "b:", b)

	a ^= b
	fmt.Println("4. a:", a, "b:", b)
}
