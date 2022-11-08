package BST

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func findLCA(p, q, root *TreeNode) *TreeNode {
	parentVal := root.Val

	pVal := p.Val
	qVal := q.Val

	if pVal > parentVal && qVal > parentVal {
		// if p and q are greater than the parent
		return findLCA(p, q, root.Right)
		// if p and q are less than the parent
	} else if pVal < parentVal && qVal < parentVal {
		return findLCA(p, q, root.Left)
	} else {
		// we found LCA node
		return root
	}
}

// iterative
func findLCA2(p, q, root *TreeNode) *TreeNode {
	node := root

	pVal := p.Val
	qVal := q.Val

	for node != nil {
		parentVal := node.Val

		if pVal > parentVal && qVal > parentVal {
			node = node.Right
		} else if pVal < parentVal && qVal < parentVal {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
