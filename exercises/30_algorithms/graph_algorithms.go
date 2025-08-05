// graph_algorithms.go
// Learn graph algorithms: BFS, DFS, shortest path, and graph traversal

package main

import (
	"fmt"
	"math"
)

// TODO: Graph representation using adjacency list
type Graph struct {
	// TODO: Define graph structure using adjacency list
}

// TODO: Graph constructor
func NewGraph(vertices int) *Graph {
	// TODO: Create new graph with given number of vertices
}

// TODO: Add edge to graph
func (g *Graph) AddEdge(from, to int, weight int) {
	// TODO: Add weighted edge between vertices
}

// TODO: Breadth-First Search
func (g *Graph) BFS(start int) []int {
	// TODO: Implement BFS traversal
	// Use queue for level-order traversal
	// Mark visited vertices
	// Return traversal order
}

// TODO: Depth-First Search
func (g *Graph) DFS(start int) []int {
	// TODO: Implement DFS traversal
	// Use stack (or recursion) for depth-first traversal
	// Mark visited vertices  
	// Return traversal order
}

// TODO: Check if path exists between two vertices
func (g *Graph) HasPath(from, to int) bool {
	// TODO: Use BFS or DFS to check connectivity
}

// TODO: Find shortest path (unweighted)
func (g *Graph) ShortestPath(from, to int) []int {
	// TODO: Use BFS to find shortest path in unweighted graph
	// Track parent vertices to reconstruct path
}

// TODO: Dijkstra's algorithm for weighted shortest path
func (g *Graph) Dijkstra(start int) ([]int, []int) {
	// TODO: Implement Dijkstra's algorithm
	// Return distances and previous vertices
	// Use priority queue for efficiency
}

// TODO: Detect cycle in directed graph
func (g *Graph) HasCycle() bool {
	// TODO: Use DFS with color coding (white, gray, black)
	// Gray vertices indicate back edge (cycle)
}

// TODO: Topological sort
func (g *Graph) TopologicalSort() []int {
	// TODO: Implement topological sort using DFS
	// Only works for DAGs (Directed Acyclic Graphs)
}

// TODO: Find strongly connected components
func (g *Graph) StronglyConnectedComponents() [][]int {
	// TODO: Use Kosaraju's algorithm
	// 1. DFS to get finish times
	// 2. Create transpose graph
	// 3. DFS on transpose in reverse finish order
}

func main() {
	fmt.Println("=== Graph Algorithms ===")
	
	fmt.Println("\n=== Graph Creation and Basic Operations ===")
	
	// TODO: Create a sample graph
	graph := /* create graph with 6 vertices */
	
	// TODO: Add edges to create a connected graph
	edges := [][]int{
		{0, 1, 4}, // vertex 0 to 1 with weight 4
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 5},
		{2, 3, 8},
		{2, 4, 10},
		{3, 4, 2},
		{3, 5, 6},
		{4, 5, 3},
	}
	
	fmt.Println("Adding edges to graph:")
	for _, edge := range edges {
		/* add edge to graph */
		fmt.Printf("  Added edge: %d -> %d (weight: %d)\n", edge[0], edge[1], edge[2])
	}
	
	fmt.Println("\n=== Graph Traversal ===")
	
	// TODO: Test BFS traversal
	fmt.Println("Breadth-First Search from vertex 0:")
	bfsResult := /* perform BFS from vertex 0 */
	fmt.Printf("  BFS order: %v\n", bfsResult)
	
	// TODO: Test DFS traversal
	fmt.Println("Depth-First Search from vertex 0:")
	dfsResult := /* perform DFS from vertex 0 */
	fmt.Printf("  DFS order: %v\n", dfsResult)
	
	fmt.Println("\n=== Path Finding ===")
	
	// TODO: Test path existence
	testPaths := [][]int{{0, 5}, {0, 3}, {1, 4}, {2, 5}}
	
	fmt.Println("Testing path existence:")
	for _, path := range testPaths {
		from, to := path[0], path[1]
		exists := /* check if path exists */
		status := "❌ No path"
		if exists {
			status = "✅ Path exists"
		}
		fmt.Printf("  %d -> %d: %s\n", from, to, status)
	}
	
	// TODO: Find shortest paths
	fmt.Println("\nShortest paths (unweighted):")
	for _, path := range testPaths {
		from, to := path[0], path[1]
		shortestPath := /* find shortest path */
		if len(shortestPath) > 0 {
			fmt.Printf("  %d -> %d: %v (length: %d)\n", from, to, shortestPath, len(shortestPath)-1)
		} else {
			fmt.Printf("  %d -> %d: No path\n", from, to)
		}
	}
	
	fmt.Println("\n=== Weighted Shortest Paths (Dijkstra) ===")
	
	// TODO: Run Dijkstra's algorithm
	startVertex := 0
	distances, previous := /* run Dijkstra from vertex 0 */
	
	fmt.Printf("Shortest distances from vertex %d:\n", startVertex)
	for i := 0; i < len(distances); i++ {
		if distances[i] == math.MaxInt {
			fmt.Printf("  To vertex %d: ∞ (unreachable)\n", i)
		} else {
			fmt.Printf("  To vertex %d: %d\n", i, distances[i])
		}
	}
	
	// TODO: Reconstruct paths using previous array
	fmt.Println("\nShortest paths with weights:")
	for i := 1; i < len(distances); i++ {
		if distances[i] != math.MaxInt {
			path := /* reconstruct path using previous array */
			fmt.Printf("  %d -> %d: %v (total weight: %d)\n", startVertex, i, path, distances[i])
		}
	}
	
	fmt.Println("\n=== Cycle Detection ===")
	
	// TODO: Test cycle detection
	fmt.Println("Testing cycle detection:")
	
	// TODO: Create graph without cycle
	acyclicGraph := /* create acyclic graph */
	/* add edges to make it acyclic */
	
	hasCycle := /* check for cycle in acyclic graph */
	fmt.Printf("  Acyclic graph has cycle: %t\n", hasCycle)
	
	// TODO: Add edge to create cycle
	/* add edge to create cycle */
	hasCycle = /* check for cycle again */
	fmt.Printf("  After adding cycle edge: %t\n", hasCycle)
	
	fmt.Println("\n=== Topological Sort ===")
	
	// TODO: Create DAG for topological sort
	dag := /* create directed acyclic graph */
	
	// TODO: Add edges representing dependencies
	dependencies := [][]int{
		{0, 1}, // task 0 must complete before task 1
		{0, 2},
		{1, 3},
		{2, 3},
		{2, 4},
		{3, 5},
		{4, 5},
	}
	
	fmt.Println("Creating dependency graph:")
	for _, dep := range dependencies {
		/* add dependency edge */
		fmt.Printf("  Task %d -> Task %d\n", dep[0], dep[1])
	}
	
	// TODO: Perform topological sort
	topoOrder := /* perform topological sort */
	fmt.Printf("Topological order (task execution order): %v\n", topoOrder)
	
	fmt.Println("\n=== Strongly Connected Components ===")
	
	// TODO: Create directed graph with SCCs
	sccGraph := /* create graph with SCCs */
	
	// TODO: Add edges to create multiple SCCs
	sccEdges := [][]int{
		{0, 1}, {1, 2}, {2, 0}, // First SCC: 0, 1, 2
		{2, 3}, {3, 4}, {4, 5}, {5, 3}, // Second SCC: 3, 4, 5
		{1, 3}, // Connection between SCCs
	}
	
	fmt.Println("Creating graph with strongly connected components:")
	for _, edge := range sccEdges {
		/* add edge for SCC graph */
		fmt.Printf("  %d -> %d\n", edge[0], edge[1])
	}
	
	// TODO: Find SCCs
	sccs := /* find strongly connected components */
	fmt.Printf("Strongly Connected Components: %v\n", sccs)
	
	fmt.Println("\n=== Graph Algorithm Applications ===")
	
	fmt.Println("Real-world applications:")
	fmt.Println("✅ BFS: Web crawling, social networks, shortest path in unweighted graphs")
	fmt.Println("✅ DFS: Maze solving, topological sorting, cycle detection")
	fmt.Println("✅ Dijkstra: GPS navigation, network routing, shortest path in weighted graphs")
	fmt.Println("✅ Topological Sort: Task scheduling, build dependencies, course prerequisites")
	fmt.Println("✅ SCC: Web page ranking, social network analysis, compiler optimization")
	
	fmt.Println("\n=== Performance Analysis ===")
	fmt.Println("Time complexities:")
	fmt.Println("  BFS/DFS: O(V + E) where V = vertices, E = edges")
	fmt.Println("  Dijkstra: O((V + E) log V) with priority queue")
	fmt.Println("  Topological Sort: O(V + E)")
	fmt.Println("  SCC (Kosaraju): O(V + E)")
}