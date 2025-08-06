// tree_traversal.go - SOLUTION
// Learn binary tree data structure and traversal algorithms

package main

import (
	"fmt"
	"math"
	"strconv"
)

// Define TreeNode struct
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Define BinaryTree struct
type BinaryTree struct {
	Root *TreeNode
}

func main() {
	fmt.Println("=== Binary Tree Traversal ===")
	
	// Create a sample binary tree
	//        4
	//       / \
	//      2   6
	//     / \ / \
	//    1  3 5  7
	tree := &BinaryTree{}
	
	// Insert nodes to build the tree above
	values := []int{4, 2, 6, 1, 3, 5, 7}
	for _, val := range values {
		tree.Insert(val)
	}
	
	fmt.Printf("Built tree with values: %v\n", values)
	
	// Demonstrate different traversal methods
	fmt.Println("\n=== Tree Traversal Methods ===")
	
	fmt.Print("In-Order (Left-Root-Right): ")
	fmt.Println(tree.InOrderTraversal())
	
	fmt.Print("Pre-Order (Root-Left-Right): ")
	fmt.Println(tree.PreOrderTraversal())
	
	fmt.Print("Post-Order (Left-Right-Root): ")
	fmt.Println(tree.PostOrderTraversal())
	
	fmt.Print("Level-Order (Breadth-First): ")
	fmt.Println(tree.LevelOrderTraversal())
	
	// Tree operations
	fmt.Println("\n=== Tree Operations ===")
	
	searchValue := 5
	found := tree.Search(searchValue)
	fmt.Printf("Search for %d: %t\n", searchValue, found)
	
	height := tree.Height()
	fmt.Printf("Tree height: %d\n", height)
	
	size := tree.Size()
	fmt.Printf("Tree size: %d nodes\n", size)
	
	minVal := tree.FindMin()
	maxVal := tree.FindMax()
	fmt.Printf("Min value: %d, Max value: %d\n", minVal, maxVal)
	
	// Tree validation
	fmt.Println("\n=== Tree Validation ===")
	
	isValid := tree.IsValidBST()
	fmt.Printf("Is valid BST: %t\n", isValid)
	
	isBalanced := tree.IsBalanced()
	fmt.Printf("Is balanced: %t\n", isBalanced)
	
	// Tree modification
	fmt.Println("\n=== Tree Modification ===")
	
	fmt.Println("Original in-order:", tree.InOrderTraversal())
	
	// Delete a node
	deleteValue := 2
	fmt.Printf("Deleting node with value %d\n", deleteValue)
	tree.Delete(deleteValue)
	
	fmt.Println("After deletion:", tree.InOrderTraversal())
	
	// Demonstrate tree rotations
	fmt.Println("\n=== Tree Rotations (for balancing) ===")
	
	// Create unbalanced tree
	unbalanced := &BinaryTree{}
	for _, val := range []int{1, 2, 3, 4, 5} {
		unbalanced.Insert(val)
	}
	
	fmt.Printf("Unbalanced tree height: %d\n", unbalanced.Height())
	fmt.Println("Unbalanced in-order:", unbalanced.InOrderTraversal())
	
	// Show how rotations can balance the tree
	fmt.Println("After right rotation on root:")
	unbalanced.Root = unbalanced.RotateRight(unbalanced.Root)
	fmt.Printf("New height: %d\n", unbalanced.Height())
	fmt.Println("After rotation:", unbalanced.InOrderTraversal())
	
	fmt.Println("\n=== Tree Structure ===")
	tree.PrintTree()
}

// Tree methods

func (bt *BinaryTree) Insert(value int) {
	bt.Root = bt.insertNode(bt.Root, value)
}

func (bt *BinaryTree) insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	
	if value < node.Value {
		node.Left = bt.insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.insertNode(node.Right, value)
	}
	// If value == node.Value, don't insert duplicates
	
	return node
}

func (bt *BinaryTree) Search(value int) bool {
	return bt.searchNode(bt.Root, value)
}

func (bt *BinaryTree) searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	
	if value == node.Value {
		return true
	} else if value < node.Value {
		return bt.searchNode(node.Left, value)
	} else {
		return bt.searchNode(node.Right, value)
	}
}

func (bt *BinaryTree) Delete(value int) {
	bt.Root = bt.deleteNode(bt.Root, value)
}

func (bt *BinaryTree) deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}
	
	if value < node.Value {
		node.Left = bt.deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.deleteNode(node.Right, value)
	} else {
		// Node to be deleted found
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		
		// Node with two children
		minRight := bt.findMin(node.Right)
		node.Value = minRight.Value
		node.Right = bt.deleteNode(node.Right, minRight.Value)
	}
	
	return node
}

func (bt *BinaryTree) findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Traversal methods

func (bt *BinaryTree) InOrderTraversal() []int {
	var result []int
	bt.inOrderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) inOrderHelper(node *TreeNode, result *[]int) {
	if node != nil {
		bt.inOrderHelper(node.Left, result)
		*result = append(*result, node.Value)
		bt.inOrderHelper(node.Right, result)
	}
}

func (bt *BinaryTree) PreOrderTraversal() []int {
	var result []int
	bt.preOrderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) preOrderHelper(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		bt.preOrderHelper(node.Left, result)
		bt.preOrderHelper(node.Right, result)
	}
}

func (bt *BinaryTree) PostOrderTraversal() []int {
	var result []int
	bt.postOrderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) postOrderHelper(node *TreeNode, result *[]int) {
	if node != nil {
		bt.postOrderHelper(node.Left, result)
		bt.postOrderHelper(node.Right, result)
		*result = append(*result, node.Value)
	}
}

func (bt *BinaryTree) LevelOrderTraversal() []int {
	if bt.Root == nil {
		return []int{}
	}
	
	var result []int
	queue := []*TreeNode{bt.Root}
	
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		
		result = append(result, node.Value)
		
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	
	return result
}

// Tree analysis methods

func (bt *BinaryTree) Height() int {
	return bt.heightHelper(bt.Root)
}

func (bt *BinaryTree) heightHelper(node *TreeNode) int {
	if node == nil {
		return -1 // Height of empty tree is -1
	}
	
	leftHeight := bt.heightHelper(node.Left)
	rightHeight := bt.heightHelper(node.Right)
	
	return 1 + max(leftHeight, rightHeight)
}

func (bt *BinaryTree) Size() int {
	return bt.sizeHelper(bt.Root)
}

func (bt *BinaryTree) sizeHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	
	return 1 + bt.sizeHelper(node.Left) + bt.sizeHelper(node.Right)
}

func (bt *BinaryTree) FindMin() int {
	if bt.Root == nil {
		return 0
	}
	
	node := bt.Root
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func (bt *BinaryTree) FindMax() int {
	if bt.Root == nil {
		return 0
	}
	
	node := bt.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Value
}

func (bt *BinaryTree) IsValidBST() bool {
	return bt.isValidBSTHelper(bt.Root, math.MinInt32, math.MaxInt32)
}

func (bt *BinaryTree) isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	
	if node.Value <= min || node.Value >= max {
		return false
	}
	
	return bt.isValidBSTHelper(node.Left, min, node.Value) &&
		bt.isValidBSTHelper(node.Right, node.Value, max)
}

func (bt *BinaryTree) IsBalanced() bool {
	balanced, _ := bt.isBalancedHelper(bt.Root)
	return balanced
}

func (bt *BinaryTree) isBalancedHelper(node *TreeNode) (bool, int) {
	if node == nil {
		return true, -1
	}
	
	leftBalanced, leftHeight := bt.isBalancedHelper(node.Left)
	rightBalanced, rightHeight := bt.isBalancedHelper(node.Right)
	
	balanced := leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1
	height := 1 + max(leftHeight, rightHeight)
	
	return balanced, height
}

// Tree rotation methods (for balancing)

func (bt *BinaryTree) RotateRight(node *TreeNode) *TreeNode {
	if node == nil || node.Left == nil {
		return node
	}
	
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	
	return newRoot
}

func (bt *BinaryTree) RotateLeft(node *TreeNode) *TreeNode {
	if node == nil || node.Right == nil {
		return node
	}
	
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	
	return newRoot
}

// Tree printing methods

func (bt *BinaryTree) PrintTree() {
	if bt.Root == nil {
		fmt.Println("Empty tree")
		return
	}
	bt.printHelper(bt.Root, "", true)
}

func (bt *BinaryTree) printHelper(node *TreeNode, prefix string, isLast bool) {
	if node == nil {
		return
	}
	
	// Print current node
	fmt.Print(prefix)
	if isLast {
		fmt.Print("└── ")
	} else {
		fmt.Print("├── ")
	}
	fmt.Println(node.Value)
	
	// Calculate prefix for children
	var childPrefix string
	if isLast {
		childPrefix = prefix + "    "
	} else {
		childPrefix = prefix + "│   "
	}
	
	// Print children
	if node.Left != nil || node.Right != nil {
		if node.Right != nil {
			bt.printHelper(node.Right, childPrefix, node.Left == nil)
		}
		if node.Left != nil {
			bt.printHelper(node.Left, childPrefix, true)
		}
	}
}

// Utility functions

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}