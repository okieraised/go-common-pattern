package main

import (
	"fmt"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

func (g *Graph) AddEdge(v, w int) {
	g.adjacencyList[v] = append(g.adjacencyList[v], w)
	g.adjacencyList[w] = append(g.adjacencyList[w], v) // For undirected graph
}

// RecursiveDFS performs DFS using recursion
func (g *Graph) RecursiveDFS(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Print(start, " ")

	for _, neighbor := range g.adjacencyList[start] {
		if !visited[neighbor] {
			g.RecursiveDFS(neighbor, visited)
		}
	}
}

// IterativeDFS performs DFS using an explicit stack
func (g *Graph) IterativeDFS(start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {
		// Pop the last element
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			fmt.Print(node, " ")

			// Push all unvisited neighbors onto the stack
			for i := len(g.adjacencyList[node]) - 1; i >= 0; i-- {
				neighbor := g.adjacencyList[node][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
}

func main() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)

	fmt.Println("Recursive DFS starting from node 0:")
	visited := make(map[int]bool)
	g.RecursiveDFS(0, visited)

	fmt.Println("\nIterative DFS starting from node 0:")
	g.IterativeDFS(0)
}
