package treesandgraphs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return helper(root, root.Val, root.Val)
}

func helper(node *TreeNode, curMax, curMin int) int {
	if node == nil {
		return curMax - curMin
	}
	curMax = max(curMax, node.Val)
	curMin = min(curMin, node.Val)
	left := helper(node.Left, curMax, curMin)
	right := helper(node.Right, curMax, curMin)
	return max(left, right)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(m int) int {
	if m < 0 {
		return m * -1
	}
	return m
}

func min(p, n int) int {
	if p < n {
		return p
	}
	return n
}

// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
