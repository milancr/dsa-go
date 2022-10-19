package treesandgraphs

func reachableNodes(n int, edges [][]int, restricted []int) int {
	// initialize count for nodes we reach
	count := 0
	// enter restricted nodes into map for quicker referencing
	resMap := map[int]bool{}
	for _, node := range restricted {
		resMap[node] = true
	}
	// create adjacency list
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	// create visited hashmap
	visited := map[int]bool{}
	// recurse
	recurse(0, graph, &count, resMap, visited)

	return count
}

func recurse(cur int, graph [][]int, count *int, restricted, visited map[int]bool) {
	// if current node is in restricted map return
	if restricted[cur] {
		return
	}
	// if current node has already been visited return
	if visited[cur] {
		return
	}
	// if current node hasnt been visited add it to hashmap
	visited[cur] = true
	// increment count since we now know its a new node
	*count++
	// call recurse on each node the current node points to in directed graph
	for _, node := range graph[cur] {
		recurse(node, graph, count, restricted, visited)
	}
}
