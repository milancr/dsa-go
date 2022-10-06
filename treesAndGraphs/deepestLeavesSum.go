package treesandgraphs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	depth := dfs(root)

	if depth == 0 {
		return root.Val
	}

	queue := []*TreeNode{root}
	for depth-1 > 0 {
		depth--
		nodesInLevel := len(queue)
		for ; nodesInLevel > 0; nodesInLevel-- {
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}
	sum := 0
	for _, node := range queue {
		sum += node.Val
	}

	return sum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(dfs(root.Left), dfs(root.Right))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// https://leetcode.com/problems/deepest-leaves-sum/
// works but not optimal, can be done with only one pass
