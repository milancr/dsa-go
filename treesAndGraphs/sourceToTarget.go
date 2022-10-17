package treesandgraphs

// direct acyclic graph
//            0     1    2    3
// graph = [[1,2], [3], [3], []]
// because its a dag we know we can only go one direction
// 0 points to 1, 2, 1 points to 3, 2 points to 3, 3 is end node

var g [][]int = [][]int{{1, 2}, {3}, {3}, {}}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}
	dfs(graph, 0, []int{}, &paths)
	return paths
}

// key here is backtracking
func dfs(graph [][]int, cur int, path []int, paths *[][]int) {
	// end is the goal node
	end := len(graph) - 1
	path = append(path, cur)
	// base case
	if cur == end { // if valid path
		// make shallow clone of path
		*paths = append(*paths, append([]int{}, path...))
		// otherwise keep searching
	} else {
		// keeps visiting each path
		for _, neighbor := range graph[cur] {
			dfs(graph, neighbor, path, paths)
		}
	}
	// when done searching path we need to pop
	path = path[:len(path)-1]
}

// track visited, check if were at end, push copy, go through all neighbors of node,
// when done were done searching path and we can remove it
// https://leetcode.com/problems/all-paths-from-source-to-target/
