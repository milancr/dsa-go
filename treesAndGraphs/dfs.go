package treesandgraphs

import "fmt"

// just for reference

// TreeNode ...
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	// logic to current node first
	// CLR
	fmt.Println(root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

// returns sorted
func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	// logic to current node in middle "in order LCR"
	fmt.Println(root.Val)
	inorder(root.Right)
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}

	postorder(root.Left)
	postorder(root.Right)
	// logic to current node last "LRC"
	fmt.Println(root.Val)
}
