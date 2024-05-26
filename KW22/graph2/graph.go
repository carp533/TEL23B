package main

import (
	"fmt"
)

// Aufgabe: programmiere die BFS (siehe verzeichnis graph)
// schreibe eine Testfunktion mit folgendem graph
// 1: {2, 4},
// 2: {3, 1},
// 3: {2},
// 4: {1, 5, 7},
// 5: {2, 4, 6, 8},
// 6: {3, 5, 9},
// 7: {4, 8},
// 8: {5, 9, 7},
// 9: {6, 8},

// Graph structure using adjacency list
type Graph struct {
	adjList map[int][]int
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// BFS performs Breadth-First Search and returns the order of nodes visited
func (g *Graph) BFS(start int) []int {
	var result []int
	visited := make(map[int]bool)
	// queue := []int{start}
	visited[start] = true

	return result
}

// DFS performs Depth-First Search and returns the order of nodes visited
func (g *Graph) DFS(start int) []int {
	var result []int
	visited := make(map[int]bool)
	g.dfsHelper(start, visited, &result)
	return result
}

// dfsHelper is a helper function for DFS
func (g *Graph) dfsHelper(node int, visited map[int]bool, result *[]int) {
	visited[node] = true
	*result = append(*result, node)

	for _, neighbor := range g.adjList[node] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited, result)
		}
	}
}

func main() {
	var numNodes, numEdges, u, v, startNode int

	fmt.Print("Number of nodes: ")
	fmt.Scan(&numNodes)

	fmt.Print("Number of edges: ")
	fmt.Scan(&numEdges)

	graph := NewGraph()

	fmt.Println("Edges:")
	for i := 0; i < numEdges; i++ {
		fmt.Scan(&u, &v)
		graph.AddEdge(u, v)
	}

	fmt.Print("Starting node: ")
	fmt.Scan(&startNode)

	traversal := graph.BFS(startNode)
	fmt.Print("BFS Traversal: ")
	for _, node := range traversal {
		fmt.Printf("%d ", node)
	}
	fmt.Println()
}
