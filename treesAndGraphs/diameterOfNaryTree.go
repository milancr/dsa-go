package treesandgraphs

// Node ...
type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	diameter := 0

	var dfs func(*Node, int) int

	dfs = func(node *Node, curDepth int) int {
		if len(node.Children) == 0 {
			return curDepth
		}

		md1, md2 := curDepth, 0

		for _, child := range node.Children {
			depth := dfs(child, curDepth+1)
			if depth > md1 {
				md2 = md1
				md1 = depth
			} else if depth > md2 {
				md2 = depth
			}
			// longest path could be calculated as the sum of the two largest depths - depth of parent node
			distance := md1 + md2 - 2*curDepth
			diameter = max(diameter, distance)
		}
		// return md1 bc it'll serve as one of two variable stores for the longest paths
		return md1
	}

	return diameter
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
