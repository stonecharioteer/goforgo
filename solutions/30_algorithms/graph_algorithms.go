// graph_algorithms.go - SOLUTION
// Learn graph algorithms: BFS, DFS, shortest path, and graph traversal

package main

import (
	"fmt"
	"math"
)

// Graph representation using adjacency list
type Graph struct {
	vertices int
	adjList  [][]Edge
}

type Edge struct {
	to     int
	weight int
}

// Graph constructor
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make([][]Edge, vertices),
	}
}

// Add edge to graph
func (g *Graph) AddEdge(from, to int, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge{to: to, weight: weight})
}

// Breadth-First Search
func (g *Graph) BFS(start int) []int {
	visited := make([]bool, g.vertices)
	result := []int{}
	queue := []int{start}
	visited[start] = true
	
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)
		
		for _, edge := range g.adjList[vertex] {
			if !visited[edge.to] {
				visited[edge.to] = true
				queue = append(queue, edge.to)
			}
		}
	}
	
	return result
}

// Depth-First Search
func (g *Graph) DFS(start int) []int {
	visited := make([]bool, g.vertices)
	result := []int{}
	g.dfsHelper(start, visited, &result)
	return result
}

func (g *Graph) dfsHelper(vertex int, visited []bool, result *[]int) {
	visited[vertex] = true
	*result = append(*result, vertex)
	
	for _, edge := range g.adjList[vertex] {
		if !visited[edge.to] {
			g.dfsHelper(edge.to, visited, result)
		}
	}
}

// Check if path exists between two vertices
func (g *Graph) HasPath(from, to int) bool {
	if from == to {
		return true
	}
	
	visited := make([]bool, g.vertices)
	queue := []int{from}
	visited[from] = true
	
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		
		for _, edge := range g.adjList[vertex] {
			if edge.to == to {
				return true
			}
			if !visited[edge.to] {
				visited[edge.to] = true
				queue = append(queue, edge.to)
			}
		}
	}
	
	return false
}

// Find shortest path (unweighted)
func (g *Graph) ShortestPath(from, to int) []int {
	if from == to {
		return []int{from}
	}
	
	visited := make([]bool, g.vertices)
	parent := make([]int, g.vertices)
	queue := []int{from}
	visited[from] = true
	
	for i := range parent {
		parent[i] = -1
	}
	
	found := false
	for len(queue) > 0 && !found {
		vertex := queue[0]
		queue = queue[1:]
		
		for _, edge := range g.adjList[vertex] {
			if edge.to == to {
				parent[to] = vertex
				found = true
				break
			}
			if !visited[edge.to] {
				visited[edge.to] = true
				parent[edge.to] = vertex
				queue = append(queue, edge.to)
			}
		}
	}
	
	if !found {
		return []int{}
	}
	
	// Reconstruct path
	path := []int{}
	current := to
	for current != -1 {
		path = append([]int{current}, path...)
		current = parent[current]
	}
	
	return path
}

// Dijkstra's algorithm for weighted shortest path
func (g *Graph) Dijkstra(start int) ([]int, []int) {
	dist := make([]int, g.vertices)
	prev := make([]int, g.vertices)
	visited := make([]bool, g.vertices)
	
	// Initialize distances
	for i := range dist {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	dist[start] = 0
	
	for {
		// Find unvisited vertex with minimum distance
		minDist := math.MaxInt
		minVertex := -1
		
		for i := 0; i < g.vertices; i++ {
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				minVertex = i
			}
		}
		
		if minVertex == -1 {
			break
		}
		
		visited[minVertex] = true
		
		// Update distances to neighbors
		for _, edge := range g.adjList[minVertex] {
			if !visited[edge.to] {
				newDist := dist[minVertex] + edge.weight
				if newDist < dist[edge.to] {
					dist[edge.to] = newDist
					prev[edge.to] = minVertex
				}
			}
		}
	}
	
	return dist, prev
}

// Detect cycle in directed graph
func (g *Graph) HasCycle() bool {
	color := make([]int, g.vertices) // 0=white, 1=gray, 2=black
	
	for i := 0; i < g.vertices; i++ {
		if color[i] == 0 {
			if g.hasCycleDFS(i, color) {
				return true
			}
		}
	}
	
	return false
}

func (g *Graph) hasCycleDFS(vertex int, color []int) bool {
	color[vertex] = 1 // Gray
	
	for _, edge := range g.adjList[vertex] {
		if color[edge.to] == 1 { // Back edge found
			return true
		}
		if color[edge.to] == 0 && g.hasCycleDFS(edge.to, color) {
			return true
		}
	}
	
	color[vertex] = 2 // Black
	return false
}

// Topological sort
func (g *Graph) TopologicalSort() []int {
	visited := make([]bool, g.vertices)
	stack := []int{}
	
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			g.topoSortDFS(i, visited, &stack)
		}
	}
	
	// Reverse stack to get topological order
	result := make([]int, len(stack))
	for i, j := 0, len(stack)-1; i < len(stack); i, j = i+1, j-1 {
		result[i] = stack[j]
	}
	
	return result
}

func (g *Graph) topoSortDFS(vertex int, visited []bool, stack *[]int) {
	visited[vertex] = true
	
	for _, edge := range g.adjList[vertex] {
		if !visited[edge.to] {
			g.topoSortDFS(edge.to, visited, stack)
		}
	}
	
	*stack = append(*stack, vertex)
}

// Find strongly connected components
func (g *Graph) StronglyConnectedComponents() [][]int {
	// Step 1: Fill vertices in stack according to their finishing times
	visited := make([]bool, g.vertices)
	stack := []int{}
	
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			g.fillOrder(i, visited, &stack)
		}
	}
	
	// Step 2: Create transpose graph
	transpose := g.getTranspose()
	
	// Step 3: Process vertices in order defined by stack
	visited = make([]bool, g.vertices)
	var sccs [][]int
	
	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if !visited[vertex] {
			var component []int
			transpose.dfsHelper(vertex, visited, &component)
			sccs = append(sccs, component)
		}
	}
	
	return sccs
}

func (g *Graph) fillOrder(vertex int, visited []bool, stack *[]int) {
	visited[vertex] = true
	
	for _, edge := range g.adjList[vertex] {
		if !visited[edge.to] {
			g.fillOrder(edge.to, visited, stack)
		}
	}
	
	*stack = append(*stack, vertex)
}

func (g *Graph) getTranspose() *Graph {
	transpose := NewGraph(g.vertices)
	
	for from := 0; from < g.vertices; from++ {
		for _, edge := range g.adjList[from] {
			transpose.AddEdge(edge.to, from, edge.weight)
		}
	}
	
	return transpose
}

func main() {
	fmt.Println("=== Graph Algorithms ===")
	
	fmt.Println("\n=== Graph Creation and Basic Operations ===")
	
	// Create a sample graph
	graph := NewGraph(6)
	
	// Add edges to create a connected graph
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
		graph.AddEdge(edge[0], edge[1], edge[2])
		fmt.Printf("  Added edge: %d -> %d (weight: %d)\n", edge[0], edge[1], edge[2])
	}
	
	fmt.Println("\n=== Graph Traversal ===")
	
	// Test BFS traversal
	fmt.Println("Breadth-First Search from vertex 0:")
	bfsResult := graph.BFS(0)
	fmt.Printf("  BFS order: %v\n", bfsResult)
	
	// Test DFS traversal
	fmt.Println("Depth-First Search from vertex 0:")
	dfsResult := graph.DFS(0)
	fmt.Printf("  DFS order: %v\n", dfsResult)
	
	fmt.Println("\n=== Path Finding ===")
	
	// Test path existence
	testPaths := [][]int{{0, 5}, {0, 3}, {1, 4}, {2, 5}}
	
	fmt.Println("Testing path existence:")
	for _, path := range testPaths {
		from, to := path[0], path[1]
		exists := graph.HasPath(from, to)
		status := "❌ No path"
		if exists {
			status = "✅ Path exists"
		}
		fmt.Printf("  %d -> %d: %s\n", from, to, status)
	}
	
	// Find shortest paths
	fmt.Println("\nShortest paths (unweighted):")
	for _, path := range testPaths {
		from, to := path[0], path[1]
		shortestPath := graph.ShortestPath(from, to)
		if len(shortestPath) > 0 {
			fmt.Printf("  %d -> %d: %v (length: %d)\n", from, to, shortestPath, len(shortestPath)-1)
		} else {
			fmt.Printf("  %d -> %d: No path\n", from, to)
		}
	}
	
	fmt.Println("\n=== Weighted Shortest Paths (Dijkstra) ===")
	
	// Run Dijkstra's algorithm
	startVertex := 0
	distances, previous := graph.Dijkstra(startVertex)
	
	fmt.Printf("Shortest distances from vertex %d:\n", startVertex)
	for i := 0; i < len(distances); i++ {
		if distances[i] == math.MaxInt {
			fmt.Printf("  To vertex %d: ∞ (unreachable)\n", i)
		} else {
			fmt.Printf("  To vertex %d: %d\n", i, distances[i])
		}
	}
	
	// Reconstruct paths using previous array
	fmt.Println("\nShortest paths with weights:")
	for i := 1; i < len(distances); i++ {
		if distances[i] != math.MaxInt {
			path := reconstructPath(previous, startVertex, i)
			fmt.Printf("  %d -> %d: %v (total weight: %d)\n", startVertex, i, path, distances[i])
		}
	}
	
	fmt.Println("\n=== Cycle Detection ===")
	
	// Test cycle detection
	fmt.Println("Testing cycle detection:")
	
	// Create graph without cycle
	acyclicGraph := NewGraph(4)
	acyclicGraph.AddEdge(0, 1, 1)
	acyclicGraph.AddEdge(1, 2, 1)
	acyclicGraph.AddEdge(0, 3, 1)
	
	hasCycle := acyclicGraph.HasCycle()
	fmt.Printf("  Acyclic graph has cycle: %t\n", hasCycle)
	
	// Add edge to create cycle
	acyclicGraph.AddEdge(2, 0, 1)
	hasCycle = acyclicGraph.HasCycle()
	fmt.Printf("  After adding cycle edge: %t\n", hasCycle)
	
	fmt.Println("\n=== Topological Sort ===")
	
	// Create DAG for topological sort
	dag := NewGraph(6)
	
	// Add edges representing dependencies
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
		dag.AddEdge(dep[0], dep[1], 1)
		fmt.Printf("  Task %d -> Task %d\n", dep[0], dep[1])
	}
	
	// Perform topological sort
	topoOrder := dag.TopologicalSort()
	fmt.Printf("Topological order (task execution order): %v\n", topoOrder)
	
	fmt.Println("\n=== Strongly Connected Components ===")
	
	// Create directed graph with SCCs
	sccGraph := NewGraph(6)
	
	// Add edges to create multiple SCCs
	sccEdges := [][]int{
		{0, 1}, {1, 2}, {2, 0}, // First SCC: 0, 1, 2
		{2, 3}, {3, 4}, {4, 5}, {5, 3}, // Second SCC: 3, 4, 5
		{1, 3}, // Connection between SCCs
	}
	
	fmt.Println("Creating graph with strongly connected components:")
	for _, edge := range sccEdges {
		sccGraph.AddEdge(edge[0], edge[1], 1)
		fmt.Printf("  %d -> %d\n", edge[0], edge[1])
	}
	
	// Find SCCs
	sccs := sccGraph.StronglyConnectedComponents()
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

// Helper function to reconstruct path from Dijkstra's previous array
func reconstructPath(previous []int, start, end int) []int {
	if previous[end] == -1 && start != end {
		return []int{}
	}
	
	path := []int{}
	current := end
	for current != -1 {
		path = append([]int{current}, path...)
		current = previous[current]
	}
	
	return path
}