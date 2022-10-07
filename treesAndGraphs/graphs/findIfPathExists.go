package graphs

// ************
// 3
// [[0,1],[1,2],[2,0]]
// 0
// 2
func validPath(n int, edges [][]int, start int, end int) bool {
	graph := make([][]int, n)
	// convert edges to adjacency list
	// [[], [], []] bc n == 3
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	// [[1 2] [0 2] [1 0]]

	visited := make(map[int]bool)
	q := []int{start}
	visited[start] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			return true
		}
		// apply bfs wt visited map
		for _, node := range graph[cur] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}

	return false
}

// explanation:

// https://leetcode.com/problems/find-if-path-exists-in-graph/
