package main

import (
	"fmt"
	"math"
)

type node struct {
	Data  int
	Left  *node
	Right *node
}

func heightByRecursion(root *node) int {
	if root == nil {
		return -1
	}

	lHeight := heightByRecursion(root.Left)
	rHeight := heightByRecursion(root.Right)

	return int(math.Max(float64(lHeight), float64(rHeight))) + 1
}

func heightByOrderTraversal(root *node) int {
	if root == nil {
		return 0
	}

	height := 0

	return height - 1
}

func main() {
	root := node{Data: 1}
	root.Left = &node{Data: 2}
	root.Right = &node{Data: 3}
	root.Left.Left = &node{Data: 4}
	root.Left.Right = &node{Data: 5}

	fmt.Println(heightByRecursion(&root))
}
