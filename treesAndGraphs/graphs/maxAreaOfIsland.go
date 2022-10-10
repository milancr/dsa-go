package graphs

// ********************
var visited [][]bool
var matrix [][]int

func maxAreaOfIsland(grid [][]int) int {
	// setup data structures
	visited = make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	matrix = grid

	maxArea := 0
	// loop through row of grid, call area function on each.
	// area function explanation below
	// if the value returned from area is greater than current maxArea, update maxArea
	for i := range grid {
		for j := range grid[i] {
			if area := area(i, j); maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

// if row is less than 0 or >= length of the matrix its out of bounds
// if col is less than 0 or >= length of the matrix its out of bounds
// if its already been visited we don't need to account for it again
// if the current value is 0 we ignore it
// if any of these cases return 0
// ELSE we add it to visited set
// do a recursive dfs on its surroundings. 1+ up + left + right + down
func area(r, c int) int {
	if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[r]) ||
		visited[r][c] || matrix[r][c] == 0 {
		return 0
	}

	visited[r][c] = true
	return 1 + area(r+1, c) + area(r-1, c) + area(r, c-1) + area(r, c+1)
}

// TC O(R*C) r = row c = col every square is visited once
// SC O(R*C) == space used by seen to keep track of visited squares and call stack from recursion
// https://www.youtube.com/watch?v=Us6LBYBVko4&ab_channel=NickWhite
// https://leetcode.com/problems/max-area-of-island/
