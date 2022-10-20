package treesandgraphs

// passes this testcase
// input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
/*
visual aid
["+","+",".","+"]
[".",".","x","+"] <- x is starting point
["+","+","+","."]
*/

// failing on this testcase and probably others
// maze: [["+","+","+"],[".",".","."],["+","+","+"]]
// entrance [1,0]
// expected 2 returning 1

func nearestExit(maze [][]byte, entrance []int) int {
	// initialize helpers for reference
	shortestDistance := 0
	height := len(maze) - 1
	width := len(maze[0]) - 1
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	queue := [][]int{entrance}

	if len(queue) > 0 {
		// pop to get current coordinates
		currentLocation := queue[len(queue)-1]
		// remove from queue
		queue = queue[:len(queue)-1]
		// separate x and y coords
		y := currentLocation[0]
		x := currentLocation[1]
		// if its not the first iteration because the you cant exit on the first move
		if len(queue) != 0 {
			// if we reach the edge return
			if y == 0 || y == height || x == 0 || x == width {
				return shortestDistance
			}
		}
		// apply all possible moves to current direction
		for _, direction := range directions {
			for i := 0; i < 2; i++ {
				currentLocation[i] = currentLocation[i] + direction[i]
			}
			y1 := currentLocation[0]
			x1 := currentLocation[1]
			// if the direction places us on a wall ignore it
			if maze[y1][x1] == '+' {
				continue
			}
			// if its a valid move append it to the queue
			queue = append(queue, currentLocation)
		}
		// if there are no valid moves return -1
		if len(queue) == 0 {
			return -1
		}
		// increment if theres a valid move, will exit at first opportunity returning the shortest path
		shortestDistance++
	}

	return shortestDistance
}
