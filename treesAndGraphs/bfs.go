package treesandgraphs

// just for reference

// TreeNode ...
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func bfs(root *TreeNode) {
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		nodesInCurrentLevel := len(queue)
		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}
