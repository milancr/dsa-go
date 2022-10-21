package treesandgraphs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	*TreeNode
	curVal int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum {
		return true
	}
	start := node{TreeNode: root, curVal: targetSum - root.Val}
	queue := []node{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.curVal == 0 {
			return true
		}

		if cur.TreeNode.Right != nil {
			queue = append(queue, node{TreeNode: cur.TreeNode.Right, curVal: targetSum - cur.TreeNode.Right.Val})
		}
		if cur.TreeNode.Left != nil {
			queue = append(queue, node{TreeNode: cur.TreeNode.Left, curVal: targetSum - cur.TreeNode.Left.Val})
		}

	}
	return false
}

func hasPathSumDFS(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum, 0)
}

func dfs(node *TreeNode, targetSum, curSum int) bool {
	if node == nil {
		return false
	}

	curSum += node.Val

	if node.Right == nil && node.Left == nil {
		return curSum == targetSum
	}

	return dfs(node.Left, targetSum, curSum) || dfs(node.Right, targetSum, curSum)
}
