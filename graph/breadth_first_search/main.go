package main

import (
	"fmt"
)

// Graph represents an adjacency list graph
type Graph struct {
	adjacencyList map[int][]int
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(v, w int) {
	g.adjacencyList[v] = append(g.adjacencyList[v], w)
	g.adjacencyList[w] = append(g.adjacencyList[w], v) // For undirected graph
}

// BFS performs Breadth-First Search starting from a given node
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool) // To track visited nodes
	queue := []int{start}         // Queue to process nodes

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0] // Dequeue the first element
		queue = queue[1:]

		fmt.Print(node, " ") // Process the node

		// Enqueue all unvisited neighbors
		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
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

	fmt.Println("BFS starting from node 0:")
	g.BFS(0)
}

// package main
//
//import (
//	"fmt"
//)
//
//// Graph represents a directed graph using an adjacency list
//type Graph struct {
//	adjacencyList map[int][]int
//}
//
//// NewGraph initializes a new directed graph
//func NewGraph() *Graph {
//	return &Graph{adjacencyList: make(map[int][]int)}
//}
//
//// AddEdge adds a **directed** edge from `v` to `w`
//func (g *Graph) AddEdge(v, w int) {
//	g.adjacencyList[v] = append(g.adjacencyList[v], w)
//}
//
//// BFS performs Breadth-First Search on a directed graph
//func (g *Graph) BFS(start int) {
//	visited := make(map[int]bool) // Track visited nodes
//	queue := []int{start}         // Queue for BFS
//
//	visited[start] = true
//
//	for len(queue) > 0 {
//		node := queue[0] // Dequeue the first element
//		queue = queue[1:]
//
//		fmt.Print(node, " ") // Process the node
//
//		// Enqueue all unvisited neighbors
//		for _, neighbor := range g.adjacencyList[node] {
//			if !visited[neighbor] {
//				visited[neighbor] = true
//				queue = append(queue, neighbor)
//			}
//		}
//	}
//}
//
//func main() {
//	g := NewGraph()
//	g.AddEdge(0, 1)
//	g.AddEdge(0, 2)
//	g.AddEdge(1, 3)
//	g.AddEdge(1, 4)
//	g.AddEdge(2, 5)
//	g.AddEdge(2, 6)
//	g.AddEdge(4, 6) // Additional directed edge
//
//	fmt.Println("BFS starting from node 0:")
//	g.BFS(0)
//}
