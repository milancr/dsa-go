package treesandgraphs

// Node ...
type Node struct {
	Val   int
	Right *Node
	Left  *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	//Initialize a queue data structure which contains just the root of the tree
	q := []*Node{root}

	for len(q) != 0 {
		size := len(q)
		i := 0
		for i < size {
			node := q[0]
			q = q[1:]
			if i < size-1 {
				node.Next = q[0]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			i++
		}
	}
	return root

}
