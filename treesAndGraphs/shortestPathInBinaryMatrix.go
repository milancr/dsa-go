package treesandgraphs

type node struct {
	x, y, dist int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid) // matrix is square so rows and cols are same
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		//if first or last grid is blocked then not possible
		return -1
	}
	queue := make([]node, 0)
	//insert first grid node with distance 1 in path queue
	queue = append(queue, node{x: 0, y: 0, dist: 1})
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		//when the last grid is found then simply return the path length
		if front.x == n-1 && front.y == n-1 {
			return front.dist
		}

		for _, direction := range directions {
			y1, x1 := direction[0]+front.y, direction[1]+front.x
			if isSafe(x1, y1, n) && grid[y1][x1] == 0 {
				queue = append(queue, node{x: x1, y: y1, dist: front.dist + 1})
				// mark visited
				grid[y1][x1] = 1
			}
		}
	}
	return -1
}

func isSafe(i, j, n int) bool {
	if i < 0 || i >= n || j < 0 || j >= n {
		return false
	}
	return true
}
