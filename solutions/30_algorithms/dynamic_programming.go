// dynamic_programming.go - SOLUTION
// Learn dynamic programming: memoization, tabulation, and optimization problems

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("=== Dynamic Programming ===")
	
	fmt.Println("\n=== Fibonacci Sequence ===")
	
	// Naive recursive fibonacci (exponential time)
	var fibonacciNaive func(int) int
	fibonacciNaive = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacciNaive(n-1) + fibonacciNaive(n-2)
	}
	
	// Memoized fibonacci (top-down DP)
	var fibonacciMemo func(int, map[int]int) int
	fibonacciMemo = func(n int, memo map[int]int) int {
		if n <= 1 {
			return n
		}
		if val, exists := memo[n]; exists {
			return val
		}
		memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
		return memo[n]
	}
	
	// Tabulated fibonacci (bottom-up DP)
	fibonacciTabulated := func(n int) int {
		if n <= 1 {
			return n
		}
		dp := make([]int, n+1)
		dp[0], dp[1] = 0, 1
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	}
	
	// Space-optimized fibonacci
	fibonacciOptimized := func(n int) int {
		if n <= 1 {
			return n
		}
		prev2, prev1 := 0, 1
		for i := 2; i <= n; i++ {
			current := prev1 + prev2
			prev2, prev1 = prev1, current
		}
		return prev1
	}
	
	fmt.Println("Computing Fibonacci numbers:")
	testValues := []int{5, 10, 15, 20}
	
	for _, n := range testValues {
		naive := fibonacciNaive(n)
		memo := fibonacciMemo(n, make(map[int]int))
		tabulated := fibonacciTabulated(n)
		optimized := fibonacciOptimized(n)
		
		fmt.Printf("  fib(%d): naive=%d, memo=%d, tab=%d, opt=%d\n", 
			n, naive, memo, tabulated, optimized)
	}
	
	fmt.Println("\n=== Longest Common Subsequence (LCS) ===")
	
	// LCS using memoization
	var lcsLength func(string, string, int, int, map[string]int) int
	lcsLength = func(s1, s2 string, i, j int, memo map[string]int) int {
		if i >= len(s1) || j >= len(s2) {
			return 0
		}
		
		key := fmt.Sprintf("%d,%d", i, j)
		if val, exists := memo[key]; exists {
			return val
		}
		
		if s1[i] == s2[j] {
			memo[key] = 1 + lcsLength(s1, s2, i+1, j+1, memo)
		} else {
			memo[key] = int(math.Max(float64(lcsLength(s1, s2, i+1, j, memo)), 
									 float64(lcsLength(s1, s2, i, j+1, memo))))
		}
		
		return memo[key]
	}
	
	// LCS using tabulation
	lcsTabulated := func(s1, s2 string) int {
		m, n := len(s1), len(s2)
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = 1 + dp[i-1][j-1]
				} else {
					dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
				}
			}
		}
		
		return dp[m][n]
	}
	
	// Reconstruct LCS string
	reconstructLCS := func(s1, s2 string, dp [][]int) string {
		var result []byte
		i, j := len(s1), len(s2)
		
		for i > 0 && j > 0 {
			if s1[i-1] == s2[j-1] {
				result = append([]byte{s1[i-1]}, result...)
				i--
				j--
			} else if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
		
		return string(result)
	}
	
	fmt.Println("Longest Common Subsequence:")
	testCases := [][]string{
		{"ABCDGH", "AEDFHR"},
		{"AGGTAB", "GXTXAYB"},
		{"programming", "logarithm"},
	}
	
	for _, testCase := range testCases {
		s1, s2 := testCase[0], testCase[1]
		
		memo := make(map[string]int)
		lengthMemo := lcsLength(s1, s2, 0, 0, memo)
		lengthTab := lcsTabulated(s1, s2)
		
		// Create DP table for reconstruction
		m, n := len(s1), len(s2)
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = 1 + dp[i-1][j-1]
				} else {
					dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
				}
			}
		}
		
		sequence := reconstructLCS(s1, s2, dp)
		
		fmt.Printf("  '%s' & '%s': length=%d (memo=%d), LCS='%s'\n", 
			s1, s2, lengthTab, lengthMemo, sequence)
	}
	
	fmt.Println("\n=== 0/1 Knapsack Problem ===")
	
	// Item structure
	type Item struct {
		Weight int
		Value  int
	}
	
	// Knapsack with memoization
	var knapsackMemo func([]Item, int, int, map[string]int) int
	knapsackMemo = func(items []Item, capacity, index int, memo map[string]int) int {
		if index >= len(items) || capacity <= 0 {
			return 0
		}
		
		key := fmt.Sprintf("%d,%d", capacity, index)
		if val, exists := memo[key]; exists {
			return val
		}
		
		// Exclude current item
		exclude := knapsackMemo(items, capacity, index+1, memo)
		
		// Include current item (if it fits)
		include := 0
		if items[index].Weight <= capacity {
			include = items[index].Value + knapsackMemo(items, capacity-items[index].Weight, index+1, memo)
		}
		
		memo[key] = int(math.Max(float64(exclude), float64(include)))
		return memo[key]
	}
	
	// Knapsack with tabulation
	knapsackTabulated := func(items []Item, capacity int) int {
		n := len(items)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, capacity+1)
		}
		
		for i := 1; i <= n; i++ {
			for w := 1; w <= capacity; w++ {
				// Exclude current item
				dp[i][w] = dp[i-1][w]
				
				// Include current item (if it fits)
				if items[i-1].Weight <= w {
					include := items[i-1].Value + dp[i-1][w-items[i-1].Weight]
					dp[i][w] = int(math.Max(float64(dp[i][w]), float64(include)))
				}
			}
		}
		
		return dp[n][capacity]
	}
	
	// Reconstruct knapsack solution
	reconstructKnapsack := func(items []Item, capacity int, dp [][]int) []Item {
		var result []Item
		n := len(items)
		w := capacity
		
		for i := n; i > 0 && w > 0; i-- {
			if dp[i][w] != dp[i-1][w] {
				result = append(result, items[i-1])
				w -= items[i-1].Weight
			}
		}
		
		return result
	}
	
	fmt.Println("0/1 Knapsack Problem:")
	
	// Create test items
	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
		{Weight: 30, Value: 120},
	}
	capacity := 50
	
	fmt.Printf("Items available (weight, value):\n")
	for i, item := range items {
		fmt.Printf("  Item %d: (%d, %d)\n", i, item.Weight, item.Value)
	}
	fmt.Printf("Knapsack capacity: %d\n", capacity)
	
	// Solve knapsack problem
	memo := make(map[string]int)
	maxValueMemo := knapsackMemo(items, capacity, 0, memo)
	maxValueTab := knapsackTabulated(items, capacity)
	
	// Create DP table for reconstruction
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			dp[i][w] = dp[i-1][w]
			if items[i-1].Weight <= w {
				include := items[i-1].Value + dp[i-1][w-items[i-1].Weight]
				dp[i][w] = int(math.Max(float64(dp[i][w]), float64(include)))
			}
		}
	}
	
	selectedItems := reconstructKnapsack(items, capacity, dp)
	
	fmt.Printf("Maximum value: %d (memo=%d)\n", maxValueTab, maxValueMemo)
	fmt.Printf("Selected items: %v\n", selectedItems)
	
	fmt.Println("\n=== Coin Change Problem ===")
	
	// Minimum coins needed (top-down)
	var coinChangeMemo func([]int, int, map[int]int) int
	coinChangeMemo = func(coins []int, amount int, memo map[int]int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		
		if val, exists := memo[amount]; exists {
			return val
		}
		
		minCoins := math.MaxInt
		for _, coin := range coins {
			result := coinChangeMemo(coins, amount-coin, memo)
			if result != -1 {
				minCoins = int(math.Min(float64(minCoins), float64(result+1)))
			}
		}
		
		if minCoins == math.MaxInt {
			memo[amount] = -1
		} else {
			memo[amount] = minCoins
		}
		
		return memo[amount]
	}
	
	// Minimum coins needed (bottom-up)
	coinChangeTabulated := func(coins []int, amount int) int {
		dp := make([]int, amount+1)
		for i := 1; i <= amount; i++ {
			dp[i] = amount + 1 // Initialize with impossible value
		}
		
		for i := 1; i <= amount; i++ {
			for _, coin := range coins {
				if coin <= i {
					dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
				}
			}
		}
		
		if dp[amount] > amount {
			return -1
		}
		return dp[amount]
	}
	
	// Count ways to make change
	coinChangeWays := func(coins []int, amount int) int {
		dp := make([]int, amount+1)
		dp[0] = 1
		
		for _, coin := range coins {
			for i := coin; i <= amount; i++ {
				dp[i] += dp[i-coin]
			}
		}
		
		return dp[amount]
	}
	
	fmt.Println("Coin Change Problem:")
	
	// Test coin change
	coins := []int{1, 3, 4}
	amounts := []int{6, 8, 11}
	
	fmt.Printf("Available coins: %v\n", coins)
	
	for _, amount := range amounts {
		memo := make(map[int]int)
		minCoinsMemo := coinChangeMemo(coins, amount, memo)
		minCoinsTab := coinChangeTabulated(coins, amount)
		ways := coinChangeWays(coins, amount)
		
		fmt.Printf("  Amount %d: min coins=%d (memo=%d), ways=%d\n", 
			amount, minCoinsTab, minCoinsMemo, ways)
	}
	
	fmt.Println("\n=== Longest Increasing Subsequence (LIS) ===")
	
	// LIS using DP
	lisLength := func(arr []int) int {
		if len(arr) == 0 {
			return 0
		}
		
		dp := make([]int, len(arr))
		for i := range dp {
			dp[i] = 1
		}
		
		for i := 1; i < len(arr); i++ {
			for j := 0; j < i; j++ {
				if arr[j] < arr[i] {
					dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				}
			}
		}
		
		maxLen := 0
		for _, length := range dp {
			maxLen = int(math.Max(float64(maxLen), float64(length)))
		}
		
		return maxLen
	}
	
	// Reconstruct LIS
	reconstructLIS := func(arr []int, dp []int) []int {
		maxLen := 0
		maxIndex := 0
		
		for i, length := range dp {
			if length > maxLen {
				maxLen = length
				maxIndex = i
			}
		}
		
		var result []int
		currentLen := maxLen
		
		for i := maxIndex; i >= 0; i-- {
			if dp[i] == currentLen {
				result = append([]int{arr[i]}, result...)
				currentLen--
			}
		}
		
		return result
	}
	
	fmt.Println("Longest Increasing Subsequence:")
	
	// Test LIS
	testArrays := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
	}
	
	for _, arr := range testArrays {
		length := lisLength(arr)
		
		// Compute DP array for reconstruction
		dp := make([]int, len(arr))
		for i := range dp {
			dp[i] = 1
		}
		for i := 1; i < len(arr); i++ {
			for j := 0; j < i; j++ {
				if arr[j] < arr[i] {
					dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				}
			}
		}
		
		sequence := reconstructLIS(arr, dp)
		
		fmt.Printf("  Array: %v\n", arr)
		fmt.Printf("  LIS length: %d, sequence: %v\n", length, sequence)
	}
	
	fmt.Println("\n=== Matrix Chain Multiplication ===")
	
	// Matrix chain multiplication
	matrixChainMultiplication := func(dimensions []int) int {
		n := len(dimensions) - 1
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
		}
		
		for length := 2; length <= n; length++ {
			for i := 0; i < n-length+1; i++ {
				j := i + length - 1
				dp[i][j] = math.MaxInt
				
				for k := i; k < j; k++ {
					cost := dp[i][k] + dp[k+1][j] + dimensions[i]*dimensions[k+1]*dimensions[j+1]
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(cost)))
				}
			}
		}
		
		return dp[0][n-1]
	}
	
	fmt.Println("Matrix Chain Multiplication:")
	
	// Test matrix chain multiplication
	dimensions := []int{40, 20, 30, 10, 30}
	fmt.Printf("Matrix dimensions: %v\n", dimensions)
	fmt.Printf("(Matrix sizes: 40x20, 20x30, 30x10, 10x30)\n")
	
	minMultiplications := matrixChainMultiplication(dimensions)
	fmt.Printf("Minimum scalar multiplications: %d\n", minMultiplications)
	
	fmt.Println("\n=== Edit Distance (Levenshtein) ===")
	
	// Edit distance using DP
	editDistance := func(s1, s2 string) int {
		m, n := len(s1), len(s2)
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		
		// Initialize base cases
		for i := 0; i <= m; i++ {
			dp[i][0] = i
		}
		for j := 0; j <= n; j++ {
			dp[0][j] = j
		}
		
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s1[i-1] == s2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = 1 + int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1])))
				}
			}
		}
		
		return dp[m][n]
	}
	
	fmt.Println("Edit Distance (Levenshtein Distance):")
	
	// Test edit distance
	stringPairs := [][]string{
		{"kitten", "sitting"},
		{"saturday", "sunday"},
		{"intention", "execution"},
	}
	
	for _, pair := range stringPairs {
		s1, s2 := pair[0], pair[1]
		distance := editDistance(s1, s2)
		fmt.Printf("  '%s' -> '%s': distance = %d\n", s1, s2, distance)
	}
	
	fmt.Println("\n=== DP Patterns and Optimizations ===")
	
	fmt.Println("Common DP patterns:")
	fmt.Println("✅ Linear DP: Fibonacci, Climbing Stairs, House Robber")
	fmt.Println("✅ 2D DP: LCS, Edit Distance, Unique Paths")
	fmt.Println("✅ Interval DP: Matrix Chain Multiplication, Palindrome Partitioning")
	fmt.Println("✅ Knapsack DP: 0/1 Knapsack, Coin Change, Subset Sum")
	fmt.Println("✅ Tree DP: Binary Tree Maximum Path Sum")
	
	fmt.Println("\nOptimization techniques:")
	fmt.Println("✅ Space optimization: Use rolling arrays when possible")
	fmt.Println("✅ Memoization vs Tabulation: Choose based on problem constraints")
	fmt.Println("✅ State compression: Reduce dimensions when possible")
	fmt.Println("✅ Early termination: Stop when optimal solution found")
}