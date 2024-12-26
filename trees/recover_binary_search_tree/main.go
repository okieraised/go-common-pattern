package main

type node struct {
	Data  int
	Left  *node
	Right *node
}

func recoverTree(root *node) {
	var prev, first, second *node
	var dfs func(*node)
	dfs = func(root *node) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if prev != nil && prev.Data > root.Data {
			if first == nil {
				first = prev
			}
			second = root
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	first.Data, second.Data = second.Data, first.Data
}

func main() {
	root := node{Data: 1}
	root.Left = &node{Data: 2}
	root.Right = &node{Data: 3}
	root.Left.Left = &node{Data: 4}
	root.Left.Right = &node{Data: 5}

	recoverTree(&root)
}
