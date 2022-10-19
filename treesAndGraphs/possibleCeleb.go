package treesandgraphs

func knows(n1, n2 int) bool {
	// function given
	return true
}

// brute force
func possibleCeleb(knows func(a int, b int) bool, n int) func(n int) int {
	return func(n int) int {

		var isCeleb func(pc, i int) bool
		isCeleb = func(pc, i int) bool {
			for j := 0; j < n; j++ {
				if j == pc {
					continue
				}
				if knows(pc, j) || !knows(j, pc) {
					return false
				}
			}
			return true
		}

		for i := 0; i < n; i++ {
			if isCeleb(i, n) {
				return i
			}
		}
		return -1

	}

}
