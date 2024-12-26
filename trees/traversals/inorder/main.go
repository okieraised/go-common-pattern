package main

import "fmt"

type node struct {
	Data  int
	Left  *node
	Right *node
}

func InorderTraversal(root *node) {
	res := make([]int, 0)

	if root == nil {
		return
	}

	res = append(res, root.Data)

	InorderTraversal(root.Left)

	fmt.Println(root.Data)

	InorderTraversal(root.Right)
}

func main() {
	root := node{Data: 1}
	root.Left = &node{Data: 2}
	root.Right = &node{Data: 3}
	root.Left.Left = &node{Data: 4}
	root.Left.Right = &node{Data: 5}
	InorderTraversal(&root)
}
