package main

import "fmt"

type node struct {
	Data  int
	Left  *node
	Right *node
}

func PreorderTraversal(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.Data)

	PreorderTraversal(root.Left)

	PreorderTraversal(root.Right)
}

func main() {
	root := node{Data: 1}
	root.Left = &node{Data: 2}
	root.Right = &node{Data: 3}
	root.Left.Left = &node{Data: 4}
	root.Left.Right = &node{Data: 5}
	PreorderTraversal(&root)
}
