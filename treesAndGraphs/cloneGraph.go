package treesandgraphs

// Node ...
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[int]*Node{}
	return dfs(visited, node)
}

func dfs(visited map[int]*Node, node *Node) *Node {
	// if theres no node return
	if node == nil {
		return nil
	}
	// check if node has been visited
	if v, ok := visited[node.Val]; ok {
		return v
	}

	// create new node
	n := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	// add to visited
	visited[node.Val] = n
	// range through current nodes neighbors, run dfs on each neighbor, append ond nodes neighbors
	// to the new nodes neighbors field
	for _, v := range node.Neighbors {
		nn := dfs(visited, v)
		n.Neighbors = append(n.Neighbors, nn)
	}
	return n
}

// https://leetcode.com/explore/featured/card/graph/619/depth-first-search-in-graph/3900/
