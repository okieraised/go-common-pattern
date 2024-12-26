package main

import "fmt"

type node struct {
	Data  int
	Left  *node
	Right *node
}

func PostorderTraversal(root *node, traversed []int) []int {

	if root == nil {
		return traversed
	}

	traversed = PostorderTraversal(root.Left, traversed)

	traversed = PostorderTraversal(root.Right, traversed)

	fmt.Println(root.Data)

	traversed = append(traversed, root.Data)

	return traversed
}

func main() {
	root := node{Data: 1}
	root.Left = &node{Data: 2}
	root.Right = &node{Data: 3}
	root.Left.Left = &node{Data: 4}
	root.Left.Right = &node{Data: 5}
	fmt.Println(PostorderTraversal(&root, []int{}))
}
