package treesandgraphs

// ******************
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// https://www.youtube.com/watch?v=RIDBLO-S7OA&ab_channel=NickWhite
// https://leetcode.com/problems/insert-into-a-binary-search-tree/
