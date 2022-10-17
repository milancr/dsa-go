package backtracking

type Set map[[2]int]bool

func cleanRoom(robot *Robot) {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(Set)
	backTrack(0, 0, 0, visited, directions, robot)
}

func goBack(r *Robot) {
	r.TurnRight()
	r.TurnRight()
	r.Move()
	r.TurnRight()
	r.TurnRight()
}

func backTrack(row int, col int, d int, visited Set, directions [][]int, robot *Robot) {
	point := [2]int{row, col}
	visited[point] = true
	robot.Clean()
	for i := 0; i < 4; i++ {
		newDirection := (d + i) % 4
		r := row + directions[newDirection][0]
		c := col + directions[newDirection][1]
		if visited[[2]int{r, c}] == false && robot.Move() {
			backTrack(r, c, newDirection, visited, directions, robot)
			goBack(robot)
		}
		robot.TurnRight()
	}
}

func buildMap() map[int][]string{
	m := map[int][]string{}
	m[2] = ["a", "b", "c"]
	m[3] = ["d", "e", "f"]
	m[4] = ["g", "h", "i"]
	m[5] = ["j", "k", "l"]
	m[6] = ["m", "n", "o"]
	m[7] = ["p", "q", "r", "s"]
	m[8] = ["t", "u", "v"]
	m[9] = ["w", "x", "y", "z"]
	return m
}

/**
 // This is the robot's control interface.
 // You should not implement it, or speculate about its implementation
* type Robot struct {
* }
*
 // Returns true if the cell in front is open and robot moves into the cell.
 // Returns false if the cell in front is blocked and robot stays in the current cell.
* func (robot *Robot) Move() bool {}
*
 // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 // Each turn will be 90 degrees.
* func (robot *Robot) TurnLeft() {}
* func (robot *Robot) TurnRight() {}
*
 // Clean the current cell.
* func (robot *Robot) Clean() {}
*/
