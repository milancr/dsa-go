package treesandgraphs

import "fmt"

func search(board [][]byte, word string, x, y int, visited map[int]struct{}) bool {
	// If we have no more letters, we've found it!
	if len(word) == 0 {
		return true
	}
	// bounds check.
	if x >= len(board) || x < 0 || y >= len(board[0]) || y < 0 {
		return false
	}
	// If we've already visited this, obviously not found.
	if _, didVisit := visited[(y<<10)|x]; didVisit {
		return false
	}
	// If the letters don't match, obviously not found.
	if board[x][y] != word[0] {
		return false
	}

	// Backup our visited data.
	visited[(y<<10)|x] = struct{}{}

	// Okay, the letters match, and we haven't visited this place.
	word = word[1:]
	for _, newX := range []int{x - 1, x + 1} {
		for _, newY := range []int{y - 1, y + 1} {
			if search(board, word, newX, newY, visited) {
				return true
			}
		}
	}
	//delete visited[(y<<10)|x]
	return false
}

func main() {
	fmt.Println("Hello, 世界")
}
