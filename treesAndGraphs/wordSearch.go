package treesandgraphs

// 33/84 test cases passed
func exist(board [][]byte, word string) bool {

	for i, row := range board {
		for j, letter := range row {
			if letter == word[0] {
				return dfs(i, j, 0, word, board)
			}
		}
	}
	return false
}

func dfs(col, row, pointer int, word string, board [][]byte) bool {
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	m := len(board)
	n := len(board[0])
	if pointer == len(word)-1 {
		return true
	}
	for _, direction := range directions {
		col = col + direction[0]
		row = row + direction[1]
		if inRange(col, row, m, n) && board[col][row] == word[pointer] {
			return dfs(col, row, pointer+1, word, board)
		}
	}
	return false
}

func inRange(h, w, m, n int) bool {
	if h < m && h >= 0 && w < n && w >= 0 {
		return true
	}
	return false
}

// https://leetcode.com/problems/word-search/
