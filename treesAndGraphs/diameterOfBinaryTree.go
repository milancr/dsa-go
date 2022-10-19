package treesandgraphs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	longestPath(root, &diameter)
	return diameter
}

func longestPath(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	// recursively find the longest path in both left and right path
	leftPath := longestPath(node.Left, diameter)
	rightPath := longestPath(node.Right, diameter)

	*diameter = max(*diameter, leftPath+rightPath)

	return max(leftPath, rightPath) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// https://leetcode.com/problems/diameter-of-binary-tree/
