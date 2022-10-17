package backtracking

// **** need more backtracking progress
func totalNQueens(n int) int {
	board := make([][]bool, n)
	for r := range board {
		board[r] = make([]bool, n)
	} // true means a queen, false an empty spot
	cols, pDiags, nDiags := make([]bool, n), make([]bool, 2*n), make([]bool, 2*n)

	res := 0
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res++
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || pDiags[row+col] || nDiags[n+row-col] {
				continue
			}

			board[row][col] = true
			cols[col] = true
			pDiags[row+col] = true
			nDiags[n+row-col] = true

			dfs(row + 1)

			board[row][col] = false
			cols[col] = false
			pDiags[row+col] = false
			nDiags[n+row-col] = false
		}
	}
	dfs(0)

	return res
}

// trick is that one increments while the other decrements
/*
	0 1 2 3
	1
	2
	3			x  <- x would be 0
*/
