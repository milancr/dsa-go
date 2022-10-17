package treesandgraphs

// should be refactored
// essentially just bfs with a flip variable and value store array but this one was fun

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	ans := [][]int{}

	if root == nil {
		return ans
	}

	queue = append(queue, root)
	ans = append(ans, []int{root.Val})

	if root.Left == nil && root.Right == nil {
		return ans
	}

	// if true go left to right, false right to left
	// because we added root.val already we went left to right over the first element, start false
	flip := false
	for len(queue) > 0 {
		levelVals := []int{}
		nodesInLevel := len(queue)
		for ; nodesInLevel > 0; nodesInLevel-- {
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				levelVals = append(levelVals, n.Left.Val)
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				levelVals = append(levelVals, n.Right.Val)
				queue = append(queue, n.Right)
			}
		}
		if flip == false {
			l, r := 0, len(levelVals)-1
			for l < r {
				levelVals[l], levelVals[r] = levelVals[r], levelVals[l]
				l++
				r--
			}
		}
		if len(levelVals) > 0 {
			ans = append(ans, levelVals)
		}
		flip = !flip
	}
	return ans
}

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal
